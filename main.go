package main

import (
	"fmt"

	"github.com/mirekwalczak/goccp/ccp"
)

func main() {
	API := ccp.API{
		AppID:  "Example_CCCP",
		Safe:   "Example_CCCP",
		Folder: "Root",
		Object: "Example_user@ccp.example.some.net",
	}
	baseURL := "abc12345.ad.some.net"
	c, err := API.GetCCP(baseURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(
		c.Content,
		c.CreationMethod,
		c.Safe,
		c.UserName,
		c.Database,
		c.PolicyID,
		c.DeviceType,
		c.Name,
		c.Folder,
		c.PasswordChangeInProcess,
		c.CPMDisabled,
	)
}
