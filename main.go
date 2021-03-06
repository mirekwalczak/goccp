package main

import (
	"fmt"

	"github.com/mirekwalczak/goccp/ccp"
)

func main() {
	API := ccp.API{
		Host:   "abc12345.ad.some.net",
		AppID:  "Example_CCCP",
		Safe:   "Example_CCCP",
		Folder: "Root",
		Object: "Example_user@ccp.example.some.net",
	}
	c, err := API.GetCCP()
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
