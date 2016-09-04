package domain

import "testing"

type mockAccessEntity struct {
	Access map[string][]string
}

func (mae mockAccessEntity) AccessTypes() map[string][]string {
	return mae.Access
}

func (mae mockAccessEntity) Owner() string {
	return "someid"
}

type mockAccessActor struct {
	ActorType string
	ID        string
}

func (maa mockAccessActor) Type() string {
	return maa.ActorType
}

func (maa mockAccessActor) Id() string {
	return maa.ID
}

func TestAuthorise(t *testing.T) {
	tests := []struct {
		ID            string
		Type          string
		ErrorExpected bool
		Access        map[string][]string
		Action        string
	}{
		//councillor accessing  wrong id
		{
			Action:        "write",
			ID:            "simple",
			Type:          "councillor",
			ErrorExpected: true,
			Access: map[string][]string{
				"read":  []string{"admin"},
				"write": []string{"admin"},
			},
		},
		//councillor accessing correct id
		{
			Action:        "read",
			ID:            "someid",
			Type:          "councillor",
			ErrorExpected: false,
			Access: map[string][]string{
				"read":  []string{"admin"},
				"write": []string{"admin"},
			},
		},
		{
			Action:        "write",
			ID:            "someid",
			Type:          "councillor",
			ErrorExpected: false,
			Access: map[string][]string{
				"read":  []string{"admin"},
				"write": []string{"admin"},
			},
		},
		//admin accessing
		{
			Action:        "write",
			ID:            "admin",
			Type:          "admin",
			ErrorExpected: false,
			Access: map[string][]string{
				"read":  []string{"admin"},
				"write": []string{"admin"},
			},
		},
		{
			Action:        "read",
			ID:            "admin",
			Type:          "admin",
			ErrorExpected: false,
			Access: map[string][]string{
				"read":  []string{"admin"},
				"write": []string{"admin"},
			},
		},
		//incorrect user
		{
			Action:        "write",
			ID:            "otherlocal",
			Type:          "local",
			ErrorExpected: true,
			Access: map[string][]string{
				"read":  []string{"admin"},
				"write": []string{"admin"},
			},
		},
		{
			Action:        "read",
			ID:            "otherlocal",
			Type:          "local",
			ErrorExpected: true,
			Access: map[string][]string{
				"read":  []string{"admin"},
				"write": []string{"admin"},
			},
		},
		//public entity all access
		{
			Action:        "read",
			ID:            "otherlocal",
			Type:          "local",
			ErrorExpected: false,
			Access: map[string][]string{
				"read":  []string{"any"},
				"write": []string{"any"},
			},
		},
		{
			Action:        "write",
			ID:            "otherlocal",
			Type:          "local",
			ErrorExpected: false,
			Access: map[string][]string{
				"read":  []string{"any"},
				"write": []string{"any"},
			},
		},
	}

	auth := &Authorisation{}
	for _, test := range tests {
		ent := &mockAccessEntity{Access: test.Access}
		act := &mockAccessActor{
			ActorType: test.Type,
			ID:        test.ID,
		}
		var err error
		if err = auth.Authorise(ent, "write", act); err != nil {
			if !test.ErrorExpected {
				t.Errorf("did not expect an error acess entity %s %v", err.Error(), test)
				continue
			}
		}
		if test.ErrorExpected && nil == err {
			t.Errorf("expected an error acessing entity but did not get one")
		}
	}
}
