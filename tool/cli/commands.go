package cli

import (
	"encoding/json"
	"fmt"
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	uuid "github.com/goadesign/goa/uuid"
	"github.com/maleck13/local/client"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type (
	// CreateCouncillorAdminCommand is the command line data structure for the createCouncillor action of admin
	CreateCouncillorAdminCommand struct {
		PrettyPrint bool
	}

	// CloseCommunicationsCommand is the command line data structure for the close action of communications
	CloseCommunicationsCommand struct {
		ID          string
		PrettyPrint bool
	}

	// ListCommunicationsCommand is the command line data structure for the list action of communications
	ListCommunicationsCommand struct {
		// recepientID
		Rid string
		// communication id
		CommID      string
		PrettyPrint bool
	}

	// RecieveEmailCommunicationsCommand is the command line data structure for the recieveEmail action of communications
	RecieveEmailCommunicationsCommand struct {
		PrettyPrint bool
	}

	// SendCommunicationsCommand is the command line data structure for the send action of communications
	SendCommunicationsCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// ListConstituentsCouncillorsCommand is the command line data structure for the listConstituents action of councillors
	ListConstituentsCouncillorsCommand struct {
		ID          string
		PrettyPrint bool
	}

	// ListForCountyAndAreaCouncillorsCommand is the command line data structure for the listForCountyAndArea action of councillors
	ListForCountyAndAreaCouncillorsCommand struct {
		// The area of the users local council
		Area string
		// The area of the users local council
		County      string
		PrettyPrint bool
	}

	// ReadByIDCouncillorsCommand is the command line data structure for the readById action of councillors
	ReadByIDCouncillorsCommand struct {
		ID          string
		PrettyPrint bool
	}

	// UpdateCouncillorsCommand is the command line data structure for the update action of councillors
	UpdateCouncillorsCommand struct {
		Payload     string
		ContentType string
		ID          string
		PrettyPrint bool
	}

	// UploadProfilePicCouncillorsCommand is the command line data structure for the uploadProfilePic action of councillors
	UploadProfilePicCouncillorsCommand struct {
		ID          string
		PrettyPrint bool
	}

	// DeleteUserCommand is the command line data structure for the delete action of user
	DeleteUserCommand struct {
		// user ID
		ID          string
		PrettyPrint bool
	}

	// ListUserCommand is the command line data structure for the list action of user
	ListUserCommand struct {
		PrettyPrint bool
	}

	// LoginUserCommand is the command line data structure for the login action of user
	LoginUserCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// ReadUserCommand is the command line data structure for the read action of user
	ReadUserCommand struct {
		// user ID
		ID          string
		PrettyPrint bool
	}

	// ResetpasswordUserCommand is the command line data structure for the resetpassword action of user
	ResetpasswordUserCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// SignUpCouncillorUserCommand is the command line data structure for the signUpCouncillor action of user
	SignUpCouncillorUserCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// SignupUserCommand is the command line data structure for the signup action of user
	SignupUserCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// UpdateUserCommand is the command line data structure for the update action of user
	UpdateUserCommand struct {
		Payload     string
		ContentType string
		// user ID
		ID          string
		PrettyPrint bool
	}

	// VerifySignupUserCommand is the command line data structure for the verifySignup action of user
	VerifySignupUserCommand struct {
		Key         string
		UID         string
		PrettyPrint bool
	}

	// DownloadCommand is the command line data structure for the download command.
	DownloadCommand struct {
		// OutFile is the path to the download output file.
		OutFile string
	}
)

// RegisterCommands registers the resource action CLI commands.
func RegisterCommands(app *cobra.Command, c *client.Client) {
	var command, sub *cobra.Command
	command = &cobra.Command{
		Use:   "close",
		Short: `recieve an email`,
	}
	tmp1 := new(CloseCommunicationsCommand)
	sub = &cobra.Command{
		Use:   `communications ["/communications/close/ID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp1.Run(c, args) },
	}
	tmp1.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp1.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "createCouncillor",
		Short: `admin api to add a councillor`,
	}
	tmp2 := new(CreateCouncillorAdminCommand)
	sub = &cobra.Command{
		Use:   `admin ["/admin/councillor"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp2.Run(c, args) },
	}
	tmp2.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp2.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "delete",
		Short: `delete a user`,
	}
	tmp3 := new(DeleteUserCommand)
	sub = &cobra.Command{
		Use:   `user ["/user/ID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp3.Run(c, args) },
	}
	tmp3.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp3.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "list",
		Short: `list action`,
	}
	tmp4 := new(ListCommunicationsCommand)
	sub = &cobra.Command{
		Use:   `communications ["/communications/councillor/RID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp4.Run(c, args) },
	}
	tmp4.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp4.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp5 := new(ListUserCommand)
	sub = &cobra.Command{
		Use:   `user ["/user"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp5.Run(c, args) },
	}
	tmp5.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp5.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "listConstituents",
		Short: `list of consituents for a councillor`,
	}
	tmp6 := new(ListConstituentsCouncillorsCommand)
	sub = &cobra.Command{
		Use:   `councillors ["/councillors/ID/consituents"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp6.Run(c, args) },
	}
	tmp6.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp6.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "listForCountyAndArea",
		Short: `list councillors based on a users details`,
	}
	tmp7 := new(ListForCountyAndAreaCouncillorsCommand)
	sub = &cobra.Command{
		Use:   `councillors ["/councillors"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp7.Run(c, args) },
	}
	tmp7.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp7.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "login",
		Short: `login user`,
	}
	tmp8 := new(LoginUserCommand)
	sub = &cobra.Command{
		Use:   `user ["/user/login"]`,
		Short: ``,
		Long: `

Payload example:

{
   "email": "Nemo itaque.",
   "signupType": "Asperiores sint deserunt magnam ut.",
   "token": "Voluptatibus molestiae accusantium error molestias est."
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp8.Run(c, args) },
	}
	tmp8.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp8.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "read",
		Short: `get a user`,
	}
	tmp9 := new(ReadUserCommand)
	sub = &cobra.Command{
		Use:   `user ["/user/ID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp9.Run(c, args) },
	}
	tmp9.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp9.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "readByID",
		Short: `read a councillor based on an id`,
	}
	tmp10 := new(ReadByIDCouncillorsCommand)
	sub = &cobra.Command{
		Use:   `councillors ["/councillors/ID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp10.Run(c, args) },
	}
	tmp10.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp10.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "recieveEmail",
		Short: `recieve an email`,
	}
	tmp11 := new(RecieveEmailCommunicationsCommand)
	sub = &cobra.Command{
		Use:   `communications ["/communications/email/recieve"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp11.Run(c, args) },
	}
	tmp11.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp11.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "resetpassword",
		Short: `resets the users password `,
	}
	tmp12 := new(ResetpasswordUserCommand)
	sub = &cobra.Command{
		Use:   `user ["/user/resetpassword"]`,
		Short: ``,
		Long: `

Payload example:

{
   "newpassword": "Quod modi eos incidunt."
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp12.Run(c, args) },
	}
	tmp12.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp12.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "send",
		Short: `send and email `,
	}
	tmp13 := new(SendCommunicationsCommand)
	sub = &cobra.Command{
		Use:   `communications ["/communications/send"]`,
		Short: ``,
		Long: `

Payload example:

{
   "body": "Sapiente ut iure ex aperiam et.",
   "commID": "Alias cumque error consequatur ducimus asperiores iste.",
   "error": "Praesentium rerum voluptatem ea iusto explicabo.",
   "from": "Distinctio similique eum et molestiae laudantium voluptatum.",
   "id": "Tempora aut aut nesciunt.",
   "isPrivate": false,
   "open": false,
   "recepientID": "Velit dolore et.",
   "sent": "1988-01-02T10:49:03Z",
   "subject": "Aut laudantium quia reiciendis.",
   "to": "Nemo similique voluptate quae ut voluptatem.",
   "type": "Optio ab iusto."
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp13.Run(c, args) },
	}
	tmp13.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp13.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "signUpCouncillor",
		Short: `handles a councillor signup. By verify the email address is a councillors email and sending out a verification email `,
	}
	tmp14 := new(SignUpCouncillorUserCommand)
	sub = &cobra.Command{
		Use:   `user ["/user/councillor/signup"]`,
		Short: ``,
		Long: `

Payload example:

{
   "email": "Asperiores et voluptas consequuntur ullam."
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp14.Run(c, args) },
	}
	tmp14.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp14.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "signup",
		Short: `Signup a user`,
	}
	tmp15 := new(SignupUserCommand)
	sub = &cobra.Command{
		Use:   `user ["/user/signup"]`,
		Short: ``,
		Long: `

Payload example:

{
   "ID": "Nostrum distinctio rerum.",
   "active": false,
   "area": "Minus quaerat ut corrupti aut omnis.",
   "county": "Culpa tempore.",
   "email": "Aut omnis tenetur.",
   "firstName": "Et amet iure omnis est voluptas.",
   "image": "Quod aut in explicabo.",
   "location": {
      "Lat": 0.3862223156607851,
      "Lon": 0.7870993871160458
   },
   "secondName": "Quia autem qui ipsum fugiat.",
   "signupType": "Praesentium in natus non asperiores accusantium.",
   "token": "Deleniti et rem.",
   "type": "Dolorem quia ipsam consequuntur maiores reprehenderit commodi."
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp15.Run(c, args) },
	}
	tmp15.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp15.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "update",
		Short: `update action`,
	}
	tmp16 := new(UpdateCouncillorsCommand)
	sub = &cobra.Command{
		Use:   `councillors ["/councillors/ID"]`,
		Short: ``,
		Long: `

Payload example:

{
   "address": "Quod cupiditate eum consequatur.",
   "area": "Cum quis modi voluptas nostrum eum doloribus.",
   "county": "Expedita itaque molestiae totam et dolores consectetur.",
   "email": "Itaque accusamus dolorem distinctio molestiae quae.",
   "facebook": "Provident voluptatem soluta et et est.",
   "firstName": "Aut aut.",
   "inOffice": false,
   "party": "Nisi impedit.",
   "phone": "Voluptate doloremque consequatur amet.",
   "secondName": "Nulla id.",
   "twitter": "Molestias reprehenderit molestias impedit sint.",
   "web": "Exercitationem nihil illum dolores voluptate."
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp16.Run(c, args) },
	}
	tmp16.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp16.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp17 := new(UpdateUserCommand)
	sub = &cobra.Command{
		Use:   `user ["/user/ID"]`,
		Short: ``,
		Long: `

Payload example:

{
   "area": "Dolorem voluptatem autem nostrum sunt.",
   "county": "Ducimus accusantium enim voluptatibus perspiciatis deleniti possimus.",
   "email": "Id sit similique autem necessitatibus ratione.",
   "firstName": "Dolores praesentium quia nisi cupiditate illo quae.",
   "id": "Ab voluptatem dolores eaque.",
   "image": "Fugit officiis veritatis autem aspernatur.",
   "secondName": "Aut ea debitis voluptatem."
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp17.Run(c, args) },
	}
	tmp17.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp17.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "uploadProfilePic",
		Short: `update a councillors details`,
	}
	tmp18 := new(UploadProfilePicCouncillorsCommand)
	sub = &cobra.Command{
		Use:   `councillors ["/councillors/ID/image"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp18.Run(c, args) },
	}
	tmp18.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp18.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "verifySignup",
		Short: `verifies a signup using a token in the  url `,
	}
	tmp19 := new(VerifySignupUserCommand)
	sub = &cobra.Command{
		Use:   `user ["/user/signup/verify"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp19.Run(c, args) },
	}
	tmp19.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp19.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)

	dl := new(DownloadCommand)
	dlc := &cobra.Command{
		Use:   "download [PATH]",
		Short: "Download file with given path",
		RunE: func(cmd *cobra.Command, args []string) error {
			return dl.Run(c, args)
		},
	}
	dlc.Flags().StringVar(&dl.OutFile, "out", "", "Output file")
	app.AddCommand(dlc)
}

func intFlagVal(name string, parsed int) *int {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func float64FlagVal(name string, parsed float64) *float64 {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func boolFlagVal(name string, parsed bool) *bool {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func stringFlagVal(name string, parsed string) *string {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func hasFlag(name string) bool {
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--"+name) {
			return true
		}
	}
	return false
}

func jsonVal(val string) (*interface{}, error) {
	var t interface{}
	err := json.Unmarshal([]byte(val), &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func jsonArray(ins []string) ([]interface{}, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []interface{}
	for _, id := range ins {
		val, err := jsonVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, val)
	}
	return vals, nil
}

func timeVal(val string) (*time.Time, error) {
	t, err := time.Parse("RFC3339", val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func timeArray(ins []string) ([]time.Time, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []time.Time
	for _, id := range ins {
		val, err := timeVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func uuidVal(val string) (*uuid.UUID, error) {
	t, err := uuid.FromString(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func uuidArray(ins []string) ([]uuid.UUID, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []uuid.UUID
	for _, id := range ins {
		val, err := uuidVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func float64Val(val string) (*float64, error) {
	t, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func float64Array(ins []string) ([]float64, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []float64
	for _, id := range ins {
		val, err := float64Val(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func boolVal(val string) (*bool, error) {
	t, err := strconv.ParseBool(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func boolArray(ins []string) ([]bool, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []bool
	for _, id := range ins {
		val, err := boolVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

// Run downloads files with given paths.
func (cmd *DownloadCommand) Run(c *client.Client, args []string) error {
	var (
		fnf func(context.Context, string) (int64, error)
		fnd func(context.Context, string, string) (int64, error)

		rpath   = args[0]
		outfile = cmd.OutFile
		logger  = goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
		ctx     = goa.WithLogger(context.Background(), logger)
		err     error
	)

	if rpath[0] != '/' {
		rpath = "/" + rpath
	}
	if rpath == "/swagger.json" {
		fnf = c.DownloadSwaggerJSON
		if outfile == "" {
			outfile = "swagger.json"
		}
		goto found
	}
	if strings.HasPrefix(rpath, "/swagger-ui/") {
		fnd = c.DownloadSwaggerUI
		rpath = rpath[12:]
		if outfile == "" {
			_, outfile = path.Split(rpath)
		}
		goto found
	}
	return fmt.Errorf("don't know how to download %s", rpath)
found:
	ctx = goa.WithLogContext(ctx, "file", outfile)
	if fnf != nil {
		_, err = fnf(ctx, outfile)
	} else {
		_, err = fnd(ctx, rpath, outfile)
	}
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	return nil
}

// Run makes the HTTP request corresponding to the CreateCouncillorAdminCommand command.
func (cmd *CreateCouncillorAdminCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/admin/councillor"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateCouncillorAdmin(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateCouncillorAdminCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the CloseCommunicationsCommand command.
func (cmd *CloseCommunicationsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/communications/close/%v", cmd.ID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CloseCommunications(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CloseCommunicationsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var id string
	cc.Flags().StringVar(&cmd.ID, "id", id, ``)
}

// Run makes the HTTP request corresponding to the ListCommunicationsCommand command.
func (cmd *ListCommunicationsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/communications/councillor/%v", cmd.Rid)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListCommunications(ctx, path, stringFlagVal("commID", cmd.CommID))
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListCommunicationsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var rid string
	cc.Flags().StringVar(&cmd.Rid, "rid", rid, `recepientID`)
	var commID string
	cc.Flags().StringVar(&cmd.CommID, "commID", commID, `communication id`)
}

// Run makes the HTTP request corresponding to the RecieveEmailCommunicationsCommand command.
func (cmd *RecieveEmailCommunicationsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/communications/email/recieve"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.RecieveEmailCommunications(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *RecieveEmailCommunicationsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the SendCommunicationsCommand command.
func (cmd *SendCommunicationsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/communications/send"
	}
	var payload client.Communication
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.SendCommunications(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *SendCommunicationsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the ListConstituentsCouncillorsCommand command.
func (cmd *ListConstituentsCouncillorsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/councillors/%v/consituents", cmd.ID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListConstituentsCouncillors(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListConstituentsCouncillorsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var id string
	cc.Flags().StringVar(&cmd.ID, "id", id, ``)
}

// Run makes the HTTP request corresponding to the ListForCountyAndAreaCouncillorsCommand command.
func (cmd *ListForCountyAndAreaCouncillorsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/councillors"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListForCountyAndAreaCouncillors(ctx, path, stringFlagVal("area", cmd.Area), stringFlagVal("county", cmd.County))
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListForCountyAndAreaCouncillorsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var area string
	cc.Flags().StringVar(&cmd.Area, "area", area, `The area of the users local council`)
	var county string
	cc.Flags().StringVar(&cmd.County, "county", county, `The area of the users local council`)
}

// Run makes the HTTP request corresponding to the ReadByIDCouncillorsCommand command.
func (cmd *ReadByIDCouncillorsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/councillors/%v", cmd.ID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ReadByIDCouncillors(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ReadByIDCouncillorsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var id string
	cc.Flags().StringVar(&cmd.ID, "id", id, ``)
}

// Run makes the HTTP request corresponding to the UpdateCouncillorsCommand command.
func (cmd *UpdateCouncillorsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/councillors/%v", cmd.ID)
	}
	var payload client.CouncillorUpdate
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UpdateCouncillors(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateCouncillorsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
	var id string
	cc.Flags().StringVar(&cmd.ID, "id", id, ``)
}

// Run makes the HTTP request corresponding to the UploadProfilePicCouncillorsCommand command.
func (cmd *UploadProfilePicCouncillorsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/councillors/%v/image", cmd.ID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UploadProfilePicCouncillors(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UploadProfilePicCouncillorsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var id string
	cc.Flags().StringVar(&cmd.ID, "id", id, ``)
}

// Run makes the HTTP request corresponding to the DeleteUserCommand command.
func (cmd *DeleteUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/user/%v", cmd.ID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.DeleteUser(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var id string
	cc.Flags().StringVar(&cmd.ID, "id", id, `user ID`)
}

// Run makes the HTTP request corresponding to the ListUserCommand command.
func (cmd *ListUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/user"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListUser(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the LoginUserCommand command.
func (cmd *LoginUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/user/login"
	}
	var payload client.Login
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.LoginUser(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *LoginUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the ReadUserCommand command.
func (cmd *ReadUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/user/%v", cmd.ID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ReadUser(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ReadUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var id string
	cc.Flags().StringVar(&cmd.ID, "id", id, `user ID`)
}

// Run makes the HTTP request corresponding to the ResetpasswordUserCommand command.
func (cmd *ResetpasswordUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/user/resetpassword"
	}
	var payload client.ResetpasswordUserPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ResetpasswordUser(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ResetpasswordUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the SignUpCouncillorUserCommand command.
func (cmd *SignUpCouncillorUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/user/councillor/signup"
	}
	var payload client.SignUpCouncillorUserPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.SignUpCouncillorUser(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *SignUpCouncillorUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the SignupUserCommand command.
func (cmd *SignupUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/user/signup"
	}
	var payload client.User
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.SignupUser(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *SignupUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the UpdateUserCommand command.
func (cmd *UpdateUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/user/%v", cmd.ID)
	}
	var payload client.UpdateUser
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UpdateUser(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
	var id string
	cc.Flags().StringVar(&cmd.ID, "id", id, `user ID`)
}

// Run makes the HTTP request corresponding to the VerifySignupUserCommand command.
func (cmd *VerifySignupUserCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/user/signup/verify"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.VerifySignupUser(ctx, path, stringFlagVal("key", cmd.Key), stringFlagVal("uid", cmd.UID))
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *VerifySignupUserCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var key string
	cc.Flags().StringVar(&cmd.Key, "key", key, ``)
	var uid string
	cc.Flags().StringVar(&cmd.UID, "uid", uid, ``)
}
