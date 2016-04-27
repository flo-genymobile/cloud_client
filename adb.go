package main

import (
    "fmt"
    "cloud_adb_client/webserver"
    "gopkg.in/gcfg.v1"
    "flag"
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

func uninstallApplicationOnVirtualMachine(id string, token string, packageName string) {
    url := webserver.GetAdbURL(id)
    adbCommandInfo := webserver.BuildAdbUninstallCommand(packageName)
    request := webserver.PrepareAdbPost(url, token, adbCommandInfo)
    webserver.DoRequest(request)
}

func stopVirtualMachine(id string, token string) {
    url := webserver.GetVirtualMachineActionURL(id, "stop")
    request := webserver.PreparePost(url, token)
    webserver.DoRequest(request)
}

func startVirtualMachine(id string, token string) {
    url := webserver.GetVirtualMachineActionURL(id, "start")
    request := webserver.PreparePost(url, token)
    webserver.DoRequest(request)
}

func main() {
    var token string
    config := loadConfiguration("/home/flo/.config/CloudClient/settings.ini")
    
    if config.UserInfo.Username == "" {
        fmt.Println("No user config found, place a valid settings.ini under ~/.config/CloudClient/")
    } else {
        token = login(config.UserInfo.Username, config.UserInfo.Password)    
    }
    
    var action string
    flag.StringVar(&action, "action", "none", "Provide an action to perform")
    
    flag.Parse()
    
    if action == "none" {
        fmt.Println("No action provided...")
    } else {
        if action == "list" {
            vms := getVirtualMachineList(token)
            webserver.PrintVirtualMachinesList(vms)
        } else if action == "install" {
            var apkPath string
            apkPath = flag.Args()[0]
            var vmId string
            vmId = flag.Args()[1]
            
            fmt.Println("Installing " + apkPath + " onto " + vmId)
            installApplicationOnVirtualMachine(vmId, token, apkPath)    
        } else if action == "uninstall" {
            var packageName string
            packageName = flag.Args()[0]
            var vmId string
            vmId = flag.Args()[1]
            
            fmt.Println("Uninstalling " + packageName + " from " + vmId)
            uninstallApplicationOnVirtualMachine(vmId, token, packageName)    
        } else if action == "push" {
            var filePath string
            filePath = flag.Args()[0]
            var vmId string
            vmId = flag.Args()[1]
            
            fmt.Println("Pushing " + filePath + " onto " + vmId)
            pushFileToVirtualMachine(vmId, token, filePath)    
        } else if action == "stop" {
            var vmId string
            vmId = flag.Args()[0]
         
            stopVirtualMachine(vmId, token)
        } else if action == "start" {
            var vmId string
            vmId = flag.Args()[0]
            
            startVirtualMachine(vmId, token)
        }
    }
}
