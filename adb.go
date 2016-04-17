package main

import (
    "fmt"
    "cloud_adb_client/webserver"
)

type AdbPushInfo struct {
    Command string `json:"cmd"`
    Arguments string `json:"args"`
    FileData string `json:"file"`
}

func login() string {
    url := webserver.GetLoginURL()
    user := webserver.User{"", ""}
    request := webserver.PrepareLogin(url, user)
    streamResponse := webserver.DoRequest(request)
    jsonResponse := webserver.ParseSigninResponse(streamResponse)
    fmt.Println(jsonResponse.Data.Token)
    
    return jsonResponse.Data.Token
}

func getVirtualMachineList(token string) {
    url := webserver.GetListInstancesURL()
    request := webserver.PrepareGet(url, token)
    streamResponse := webserver.DoRequest(request)
   
    virtualMachineList := webserver.ParseGetInstancesResponse(streamResponse)  
    fmt.Println("Found instance 0: ", virtualMachineList.Data)
}

func main() {
    token := login()
    getVirtualMachineList(token)
    
    //Todo: send adb command
    
    fmt.Printf("done\n")
}
