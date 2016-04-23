package main

import (
    "fmt"
    "cloud_adb_client/webserver"
    "gopkg.in/gcfg.v1"
)

type Config struct {
    UserInfo struct {
        Username string
        Password string    
    }
}

func loadConfiguration(configFilePath string) Config {
    var config Config
    error := gcfg.ReadFileInto(&config, configFilePath)
    
    if error != nil {
        fmt.Println("Error while reading config file", error)
    }
    
    return config
}

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
    config := loadConfiguration("/home/flo/.config/CloudClient/settings.ini")
    
    if config.UserInfo.Username == "" {
        fmt.Println("No user config found, place a valid settings.ini under ~/.config/CloudClient/")
    } else {
        login(config.UserInfo.Username, config.UserInfo.Password)    
    }
    
    /*vms := getVirtualMachineList(token)
    webserver.PrintVirtualMachinesList(vms)
    if len(vms) > 0 {
        pushFileToVirtualMachine(vms[0].Id, token, "/home/flo/file.txt")
        installApplicationOnVirtualMachine(vms[0].Id, token, "/home/flo/Downloads/FDroid.apk")    
    }*/
}
