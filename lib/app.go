package vmware_method_operation

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/vmware/govmomi/vim25/types"

	"github.com/vmware/govmomi/object"

	"github.com/pkg/errors"

	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25/mo"

	"github.com/vmware/govmomi"

	"github.com/vmware/govmomi/vim25/soap"

	"github.com/urfave/cli"
)

func exit(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	os.Exit(1)
}

func generate_client(c *cli.Context) *govmomi.Client {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if c.String("url") == "" {
		exit(errors.New("Specify vCenter URL"))
	}

	u, err := soap.ParseURL(c.String("url"))
	if err != nil {
		exit(err)
	}

	u.User = url.UserPassword(c.String("username"), c.String("password"))

	client, err := govmomi.NewClient(ctx, u, c.Bool("insecure"))
	if err != nil {
		exit(err)
	}

	return client
}

func get_vm(c *cli.Context) []mo.VirtualMachine {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client := generate_client(c)

	m := view.NewManager(client.Client)

	v, _ := m.CreateContainerView(ctx, client.ServiceContent.RootFolder, []string{"VirtualMachine"}, c.Bool("insecure"))

	var vms []mo.VirtualMachine
	_ = v.Retrieve(ctx, []string{"VirtualMachine"}, []string{"name", "disabledMethod"}, &vms)

	var vm_obj []mo.VirtualMachine
	for _, vm := range vms {
		if vm.Name == c.String("vm") {
			vm_obj = append(vm_obj, vm)
			break
		}
	}

	return vm_obj
}

func get_disable_methods(c *cli.Context) {
	vm_obj := get_vm(c)

	if len(vm_obj) != 0 {
		for _, method := range vm_obj[0].DisabledMethod {
			fmt.Println(method)
		}
	} else {
		exit(errors.New("Not found " + c.String("vm")))
	}
}

func enable_methods(c *cli.Context) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client := generate_client(c)
	vm_obj := get_vm(c)

	o := object.NewAuthorizationManager(client.Client)
	err := o.EnableMethods(ctx, []types.ManagedObjectReference{vm_obj[0].Reference()}, c.StringSlice("list"), "abracadabra")

	if err != nil {
		exit(err)
	}
}

func disable_methods(c *cli.Context) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client := generate_client(c)
	vm_obj := get_vm(c)

	o := object.NewAuthorizationManager(client.Client)

	disable_methods := []object.DisabledMethodRequest{}
	for _, method := range c.StringSlice("list") {
		disable_methods = append(disable_methods, object.DisabledMethodRequest{
			Method: method,
		})
	}

	err := o.DisableMethods(ctx, []types.ManagedObjectReference{vm_obj[0].Reference()}, disable_methods, "abracadabra")

	if err != nil {
		exit(err)
	}
}

func Do() {
	app := cli.NewApp()
	app.Name = "vmware-method-operation"
	app.Usage = ""
	app.Commands = commands
	app.Run(os.Args)
}
