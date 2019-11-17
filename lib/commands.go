package vmware_method_operation

import (
	"github.com/urfave/cli"
)

var commands = []cli.Command{
	Method,
}

var Method = cli.Command{
	Name:      "method",
	ShortName: "m",
	Usage:     "Command to enable/disable method.",
	Subcommands: cli.Commands{
		cli.Command{
			Name:      "list",
			ShortName: "l",
			Usage:     "Display disabled methods.",
			Flags: []cli.Flag{
				URL,
				UserName,
				Password,
				Insecure,
				VM,
			},
			Action: func(c *cli.Context) {
				get_disable_methods(c)
			},
		},
		cli.Command{
			Name:      "enable",
			ShortName: "e",
			Usage:     "Enable methods.",
			Flags: []cli.Flag{
				URL,
				UserName,
				Password,
				Insecure,
				VM,
				List,
			},
			Action: func(c *cli.Context) {
				enable_methods(c)
			},
		},
		cli.Command{
			Name:      "disable",
			ShortName: "d",
			Usage:     "Disable methods.",
			Flags: []cli.Flag{
				URL,
				UserName,
				Password,
				Insecure,
				VM,
				List,
			},
			Action: func(c *cli.Context) {
				disable_methods(c)
			},
		},
	},
}
