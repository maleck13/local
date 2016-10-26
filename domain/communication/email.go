package communication

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/maleck13/local/config"
	"github.com/pkg/errors"
	"github.com/sendgrid/rest"
	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type email struct {
	Config    *config.Config
	Transport func(request rest.Request) (*rest.Response, error)
}

func (email) CommunicationType() string {
	return "email"
}

func (e email) Template(name string) (*template.Template, error) {
	var t = template.New("email")
	var err error
	if name == "communication" {
		t, err = t.Parse(EmailCommunicationTemplate)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse email template")
		}

	} else if name == "verification" {
		t, err = t.Parse(EmailVerificationTemplate)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse email template")
		}
	} else {
		return nil, errors.New("unknown template " + name)
	}
	return t, nil

}

// Communicate implements the Communicator inteface. It nows how to Communicate via email
func (e email) Communicate(sender Sender, reciever Reciever, message Message) error {
	if !e.Config.Sendgrid.Enabled {
		return errors.New("email sending is disabled")
	}
	if e.Transport == nil {
		e.Transport = sendgrid.API
	}
	from := mail.NewEmail("Locals.ie", "communications@email.locals.ie")
	subject := message.Subject()
	to := mail.NewEmail(reciever.Name(), reciever.Address())
	templated, err := e.Template(message.Template())
	if err != nil {
		return err
	}
	var b bytes.Buffer
	if err := templated.Execute(&b, struct {
		Message  Message
		Sender   Sender
		Reciever Reciever
	}{
		Message:  message,
		Sender:   sender,
		Reciever: reciever,
	}); err != nil {
		return errors.Wrap(err, "failed to Execute template")
	}
	content := mail.NewContent("text/plain", b.String())
	m := mail.NewV3MailInit(from, subject, to, content)
	if reciever.Address() != sender.Address() {
		p := &mail.Personalization{
			To:      []*mail.Email{mail.NewEmail(reciever.Name(), reciever.Address())},
			Headers: map[string]string{"commid": "test"},
			CC:      []*mail.Email{mail.NewEmail(sender.Name(), sender.Address())},
		}
		m.AddPersonalizations(p)
	}
	request := sendgrid.GetRequest(e.Config.Sendgrid.APIKey, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	response, err := e.Transport(request)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to send mail via sendgrid response status %d", response.StatusCode))
	}
	if response.StatusCode >= 300 {
		return errors.New(fmt.Sprintf("failed to send mail. status from provider %d. Body  %s ", response.StatusCode, response.Body))
	}
	return nil
}

var EmailCommunicationTemplate = `
Dear {{.Reciever.Name}} ,

This mail has been sent on behalf of {{.Sender.Name}} via locals.ie . You can reply directly to this message
and {{.Sender.Name}} will recieve it. You can also reply via the locals.ie site at {{.Message.Href}}. 
Below is the full message from {{.Sender.Name}} : 

{{.Message.Subject}}

{{.Message.Body}}


`

var EmailVerificationTemplate = `
Thank you for signing up to Locals.ie. Please click the verification link below to complete your signup.

{{.Message.Body}}

Regards,

Locals.ie

`

// NewEmailer returns an emailSender that implements Sender
func NewEmailer(name, email string) emailer {
	return emailer{
		name:  name,
		email: email,
	}
}

type emailer struct {
	email, name string
}

func (ec emailer) Address() string {
	return ec.email
}
func (ec emailer) Name() string {
	return ec.name
}

// NewEmailMessage represents an email message
func NewEmailMessage(subject, body, href, templateName string) message {
	return message{
		subject:  subject,
		body:     body,
		href:     href,
		template: templateName,
	}
}

type message struct {
	subject, body, href, template string
}

func (m message) Body() string {
	return m.body
}

func (m message) Subject() string {
	return m.subject
}
func (m message) Href() string {
	return m.href
}
func (m message) Template() string {
	return m.template
}
