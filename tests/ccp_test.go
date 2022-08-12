package tests

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/mirekwalczak/goccp/ccp"
)

func TestGetCCPOK(t *testing.T) {
	t.Parallel()
	var api ccp.API
	body := `{
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
	}`
	expected := &ccp.CentralCredentialProvider{
		Content:                 "ThisIsAnExample123",
		CreationMethod:          "PVWA",
		Address:                 "ccp.example.some.net",
		Safe:                    "Example_CCCP",
		UserName:                "Example_user",
		Database:                "Example_DB",
		PolicyID:                "GLB0005_Generic_unmanaged-D1",
		DeviceType:              "Application",
		Name:                    "Example_user@ccp.example.some.net",
		Folder:                  "Root",
		PasswordChangeInProcess: "False",
	}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, body)
	})
	srv := httptest.NewServer(handler)
	defer srv.Close()
	res, err := api.GetCCP(srv.URL)
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	if res != nil && !reflect.DeepEqual(res, expected) {
		t.Errorf("expected res to be %v got %v", expected, res)
	}
}

func TestGetCCPInvalidData(t *testing.T) {
	t.Parallel()
	var api ccp.API
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"name":"abc"}`)
	})
	srv := httptest.NewServer(handler)
	defer srv.Close()
	_, err := api.GetCCP(srv.URL)
	var invalidData *ccp.InvalidDataError
	if !errors.As(err, &invalidData) {
		t.Errorf("expected err *ccp.InvalidDataError but got %T", err)
	}
}

func TestGetCCPInvalidJSON(t *testing.T) {
	t.Parallel()
	var api ccp.API
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "abc")
	})
	srv := httptest.NewServer(handler)
	defer srv.Close()
	_, err := api.GetCCP(srv.URL)
	var invalidJson *json.SyntaxError
	if !errors.As(err, &invalidJson) {
		t.Errorf("expected err *json.SyntaxError but got %T", err)
	}
}
