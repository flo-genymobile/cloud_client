package main

import (
    "fmt"
    "cloud_adb_client/webserver"
)

func login() string {
    url := webserver.GetLoginURL()
    user := webserver.User{"", ""}
    request := webserver.PrepareLogin(url, user)
    streamResponse := webserver.DoRequest(request)
    jsonResponse := webserver.ParseSigninResponse(streamResponse)
    fmt.Println(jsonResponse.Data.Token)
    
    return jsonResponse.Data.Token
}

func getVirtualMachineList(token string) []webserver.VirtualMachine {
    url := webserver.GetListInstancesURL()
    request := webserver.PrepareGet(url, token)
    streamResponse := webserver.DoRequest(request)
   
    virtualMachineList := webserver.ParseGetInstancesResponse(streamResponse)  
    fmt.Println("Found instance 0: ", virtualMachineList.Data)
    
    return virtualMachineList.Data.VirtualMachines
}

func pushFileToVirtualMachine(id string, token string, filePath string) {
    url := webserver.GetAdbURL(id)
    adbCommandInfo := webserver.BuildAdbPushCommand(filePath)
    request := webserver.PrepareAdbPost(url, token, adbCommandInfo)
    webserver.DoRequest(request)
}

func installApplicationOnVirtualMachine(id string, token string, apkPath string) {
    url := webserver.GetAdbURL(id)
    adbCommandInfo := webserver.BuildAdbInstallCommand(apkPath)
    request := webserver.PrepareAdbPost(url, token, adbCommandInfo)
    webserver.DoRequest(request)
}

func main() {
    token := login()
    vms := getVirtualMachineList(token)
    webserver.PrintVirtualMachinesList(vms)
    pushFileToVirtualMachine(vms[0].Id, token, "/home/flo/file.txt")
    installApplicationOnVirtualMachine(vms[0].Id, token, "/home/flo/Downloads/FDroid.apk")
}
