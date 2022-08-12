package ccp

import (
	"errors"
	"testing"
)

func TestToCCP(t *testing.T) {
	t.Parallel()
	httpBody := []byte(`
        {
            "Content": "ThisIsAnExample123",
            "CreationMethod": "PVWA",
            "Address": "ccp.example.some.net",
            "Safe": "Example_CCCP",
            "UserName": "Example_user",
            "Database": "Example_DB",
            "PolicyID": "GLB0005_Generic_unmanaged-D1",
            "DeviceType": "Application",
            "Name": "Example_user@ccp.example.some.net",
            "Folder": "Root",
            "PasswordChangeInProcess": "False"
        }
    `)
	c, err := toCCP(httpBody)
	if err != nil {
		t.Fatalf("toCCP failed: %s", err)
	}

	var tests = []struct {
		input string
		want  string
	}{
		{c.Content, "ThisIsAnExample123"},
		{c.CreationMethod, "PVWA"},
		{c.Address, "ccp.example.some.net"},
		{c.Safe, "Example_CCCP"},
		{c.UserName, "Example_user"},
		{c.Database, "Example_DB"},
		{c.PolicyID, "GLB0005_Generic_unmanaged-D1"},
		{c.DeviceType, "Application"},
		{c.Name, "Example_user@ccp.example.some.net"},
		{c.Folder, "Root"},
		{c.PasswordChangeInProcess, "False"},
	}

	for _, test := range tests {
		if test.input != test.want {
			t.Errorf("Should be (%v) but is (%v)", test.want, test.input)
		}
	}
}

func TestToError(t *testing.T) {
	t.Parallel()
	httpBody := []byte(`
		{
			"ErrorCode": "APPAP008E",
			"ErrorMsg": "Problem occurred while trying to use user in the Vault"
		}`)

	if err := toError(httpBody); err != nil {
		var e *InvalidDataError
		if errors.As(err, &e) {
			if e.ErrorCode != "APPAP008E" {
				t.Errorf("Should be (%v) but is (%v)", "APPAP008E", e.ErrorCode)
			}
			if e.ErrorMsg != "Problem occurred while trying to use user in the Vault" {
				t.Errorf("Should be (%v) but is (%v)",
					"Problem occurred while trying to use user in the Vault",
					e.ErrorMsg)
			}
		} else {
			t.Fatalf("toError failed: %s", err)
		}
	}
}
