package webserver

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