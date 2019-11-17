package vmware_method_operation

import (
	"github.com/urfave/cli"
)

var flags = []cli.Flag{
	URL,
	UserName,
	Password,
	Insecure,
}

var URL = cli.StringFlag{
	Name:  "url, u",
	Usage: "ESX or vCenter URL",
}

var UserName = cli.StringFlag{
	Name:  "username",
	Usage: "vCenter User Name",
	Value: "administrator@vsphere.local",
}

var Password = cli.StringFlag{
	Name:  "password",
	Usage: "vCenter User Password",
}

var Insecure = cli.BoolFlag{
	Name:  "insecure",
	Usage: "Validate flag for server certificate.",
}

var VM = cli.StringFlag{
	Name:  "vm",
	Usage: "VM name to enable/disable the method.",
}

var List = cli.StringSliceFlag{
	Name:  "list, l",
	Usage: "Method list.",
}
