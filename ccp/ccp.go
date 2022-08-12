package ccp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// API represents request data for Central Credential Provider
type API struct {
	AppID  string
	Safe   string
	Folder string
	Object string
}

// GetCCP returns response from Central Credential Provider API
func (r *API) GetCCP(baseURL string) (*CentralCredentialProvider, error) {
	url := fmt.Sprintf("%s/AIMWebService/api/Accounts?AppID=%s&Safe=%s&Folder=%s&Object=%s",
	baseURL, r.AppID, r.Safe, r.Folder, r.Object)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		if err = toError(body); err != nil {
			return nil, err
		}
	}
	ccp, err := toCCP(body)
	if err != nil {
		return nil, err
	}

	return ccp, nil
}

// CentralCredentialProvider represents Central Credential Provider response object
type CentralCredentialProvider struct {
	Content                 string `json:"Content"`
	CreationMethod          string `json:"CreationMethod"`
	Address                 string `json:"Address"`
	Safe                    string `json:"Safe"`
	UserName                string `json:"UserName"`
	Database                string `json:"Database"`
	PolicyID                string `json:"PolicyID"`
	DeviceType              string `json:"DeviceType"`
	Name                    string `json:"Name"`
	Folder                  string `json:"Folder"`
	PasswordChangeInProcess string `json:"PasswordChangeInProcess"`
	CPMDisabled             string `json:"CPMDisabled"`
}

// InvalidDataError represents Central Credential Provider error response object
type InvalidDataError struct {
	ErrorCode string `json:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg"`
}

func (e *InvalidDataError) Error() string {
	return fmt.Sprintf("%s:%s: invalid data error", e.ErrorCode, e.ErrorMsg)
}

func toCCP(data []byte) (*CentralCredentialProvider, error) {
	var c *CentralCredentialProvider
	err := json.Unmarshal(data, &c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func toError(data []byte) error {
	var e *InvalidDataError
	err := json.Unmarshal(data, &e)
	if err != nil {
		return err
	}
	return &InvalidDataError{e.ErrorCode, e.ErrorMsg}
}
