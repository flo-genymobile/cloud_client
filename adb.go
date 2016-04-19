package main

import (
    "fmt"
    "cloud_adb_client/webserver"
    "flag"
)

func login(username string, password string) string {
    url := webserver.GetLoginURL()
    user := webserver.User{username, password}
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
    userPointer := flag.String("user", "foo", "valid genymotion cloud username")
    passwordPointer := flag.String("password", "foo", "valid genymotion cloud password")
    flag.Parse()
    
    token := login(*userPointer, *passwordPointer)
    vms := getVirtualMachineList(token)
    webserver.PrintVirtualMachinesList(vms)
    if len(vms) > 0 {
        pushFileToVirtualMachine(vms[0].Id, token, "/home/flo/file.txt")
        installApplicationOnVirtualMachine(vms[0].Id, token, "/home/flo/Downloads/FDroid.apk")    
    }
}
