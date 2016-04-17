package main

import (
    "fmt"
    "net/http"
    "bytes"
    "encoding/json"
    "io"
)

type SigningData struct {
    Message string `json:"message"`
    Token string `json:"token"`
}

type SigninSucess struct {
    Data SigningData `json:"data"`
    Version float64 `json:"version"`
}

type LoginInfo struct {
    Email string `json:"email"`
    Password string `json:"password"`
}

type InstancesInfo struct {
    Data Instances `json:"data"`  
    Version float64 `json:"version"`
}

type Instances struct {
    InstanceList []Instance `json:"instances"`
}

type Instance struct {
    GenyId string `json:"genyid"`
    Hostname string `json:"hostname"`
    Id string `json:"id"`
    InstanceState string `json:"instance_state"`
    Name string `json:"name"`
    Port int `json:"port"`
    Server int `json:"server"`
    Uuid string `json:"uuid"`
}

type AdbPushInfo struct {
    Command string `json:"cmd"`
    Arguments string `json:"args"`
    FileData string `json:"file"`
}

func prepareRequest(url string, loginInfo LoginInfo) *http.Request {
    payload, err := json.Marshal(loginInfo)
    if err != nil {
        fmt.Println("can't marhsall: ", err)
    }
    
    request, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
    if err != nil {
        fmt.Println("can't create request: ", err)
    }
    request.Header.Set("Content-Type", "application/json")
    
    return request
}

func parseJsonResponse(inputStream io.ReadCloser) SigninSucess {
    decoder := json.NewDecoder(inputStream)
    var response SigninSucess
    
    var err = decoder.Decode(&response)
    if err != nil {
        fmt.Println("%T\n%s\n%#v\n",err, err, err)
    }
    
    return response
}

func main() {
    baseUrl := "https://api-cloud.genymotion.com/v1"
    authenticationEndpoint := "/users/signin"
    listInstancesEndpoint := "/instances"
    url := baseUrl + authenticationEndpoint

    //build request
    loginInfo := LoginInfo{"", ""}
    request := prepareRequest(url, loginInfo)
    
    //post call to endpoint
    httpClient := &http.Client{}
    httpresponse, err := httpClient.Do(request)
    if err != nil {
        fmt.Println("POST failed: ", err)
    }
    defer httpresponse.Body.Close()
    
    //parse server response and show token
    response := parseJsonResponse(httpresponse.Body)
    fmt.Println(response.Data.Token)
    
    //TODO: extract into functions
    //list available instances
    getInstanceUrl := baseUrl + listInstancesEndpoint
    getInstanceRequest, error := http.NewRequest("GET", getInstanceUrl, nil)
    getInstanceRequest.Header.Set("Authorization", "Bearer " + response.Data.Token)
    if error != nil {
        fmt.Println("can't create request: ", error)
    }
    
    getInstanceResponse, getInstanceErr := httpClient.Do(getInstanceRequest)
    if getInstanceErr != nil {
        fmt.Println("GET failed: ", getInstanceErr)
    }
    defer getInstanceResponse.Body.Close()
    fmt.Println("Req Status Code: ", getInstanceResponse.StatusCode) //TODO: handle errors
    
    decoder := json.NewDecoder(getInstanceResponse.Body)
    var instanceResponse InstancesInfo
    
    var decodeInstanceErr = decoder.Decode(&instanceResponse)
    if decodeInstanceErr != nil {
        fmt.Println("%T\n%s\n%#v\n", decodeInstanceErr, decodeInstanceErr, decodeInstanceErr)
    }
    
    fmt.Println("Found instance 0: ", instanceResponse.Data)
    
    //Send adb command
    
    
    fmt.Printf("done\n")
}
