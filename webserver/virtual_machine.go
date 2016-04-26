package webserver

import (
    "fmt"
)

type InstancesInfo struct {
    Data Instances `json:"data"`  
    Version float64 `json:"version"`
}

type Instances struct {
    VirtualMachines []VirtualMachine `json:"instances"`
}

type VirtualMachine struct {
    GenyId string `json:"genyid"`
    Hostname string `json:"hostname"`
    Id string `json:"id"`
    InstanceState string `json:"instance_state"`
    Name string `json:"name"`
    Port int `json:"port"`
    Server int `json:"server"`
    Uuid string `json:"uuid"`
}

func PrintVirtualMachinesList(virtualMachines []VirtualMachine) {
    if virtualMachines != nil && len(virtualMachines) > 0 {
        fmt.Println("|---------------ID-------------------||------STATE------||----------------NAME-----------------------------------|")
        for index := 0; index < len(virtualMachines); index++ {
            printVirtualMachine(virtualMachines[index])
        }
    } else {
        fmt.Println("No running Virtual Machines were found....")
    }
}

func printVirtualMachine(virtualMachine VirtualMachine) {
    fmt.Printf("|%v|", virtualMachine.Id)
    fmt.Printf("|    %v      |", virtualMachine.InstanceState)
    fmt.Printf("|    %v", virtualMachine.Name)
    fmt.Printf("\n")
}