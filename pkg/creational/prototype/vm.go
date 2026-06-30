package prototype

import "fmt"

type VirtualMachine struct {
	CPU      int
	MemoryGB int
	OS       string
	Tags     map[string]string
}

// deep clone
func (vm *VirtualMachine) Clone() *VirtualMachine {
	if vm == nil {
		return nil
	}

	newVM := &VirtualMachine{
		CPU:      vm.CPU,
		MemoryGB: vm.MemoryGB,
		OS:       vm.OS,
	}

	if vm.Tags != nil {
		newVM.Tags = make(map[string]string, len(vm.Tags))
		for k, v := range vm.Tags {
			newVM.Tags[k] = v
		}
	}

	return newVM
}

func (vm *VirtualMachine) String() string {
	return fmt.Sprintf("VM [CPU: %d, RAM: %dGB, OS: %s, Tags: %v]", vm.CPU, vm.MemoryGB, vm.OS, vm.Tags)
}
