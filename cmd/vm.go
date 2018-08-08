package cmd

import (
	"github.com/exoscale/egoscale"
	"github.com/spf13/cobra"
)

// vmCmd represents the vm command
var vmCmd = &cobra.Command{
	Use:   "vm",
	Short: "Virtual machines management",
}

func getVMWithNameOrID(name string) (*egoscale.VirtualMachine, error) {
	vm := &egoscale.VirtualMachine{ID: name}
	if err := cs.GetWithContext(gContext, vm); err == nil {
		return vm, err
	}

	vm.Name = name
	vm.ID = ""

	if err := cs.GetWithContext(gContext, vm); err != nil {
		return nil, err
	}
	return vm, nil
}

func getSecurityGroup(vm *egoscale.VirtualMachine) []string {
	sgs := []string{}
	for _, sgN := range vm.SecurityGroup {
		sgs = append(sgs, sgN.Name)
	}
	return sgs
}

func init() {
	RootCmd.AddCommand(vmCmd)
}
