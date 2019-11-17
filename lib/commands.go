package vmware_method_operation

import (
	"github.com/urfave/cli"
)

var commands = []*cli.Command{
	Method,
}

var Method = &cli.Command{
	Name:    "method",
	Aliases: []string{"m"},
	Usage:   "Command to enable/disable method.",
	Subcommands: []*cli.Command{
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "Display disabled methods.",
			Flags: []cli.Flag{
				URL,
				UserName,
				Password,
				Insecure,
				VM,
			},
			Action: func(c *cli.Context) error {
				get_disable_methods(c)
				return nil
			},
		},
		{
			Name:    "enable",
			Aliases: []string{"e"},
			Usage:   "Enable methods.",
			Flags: []cli.Flag{
				URL,
				UserName,
				Password,
				Insecure,
				VM,
				List,
			},
			Action: func(c *cli.Context) error {
				enable_methods(c)
				return nil
			},
		},
		{
			Name:    "disable",
			Aliases: []string{"d"},
			Usage:   "Disable methods.",
			Flags: []cli.Flag{
				URL,
				UserName,
				Password,
				Insecure,
				VM,
				List,
			},
			Action: func(c *cli.Context) error {
				disable_methods(c)
				return nil
			},
		},
	},
}
