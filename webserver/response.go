package webserver

import (
    "fmt"
    "encoding/json"
    "io"
)

func ParseSigninResponse(inputStream io.ReadCloser) SigninResponse {
    decoder := json.NewDecoder(inputStream)
    var response SigninResponse
    
    error := decoder.Decode(&response)
    if error != nil {
        fmt.Println("%T\n%s\n%#v\n", error, error, error)
    }
    
    return response
}

func ParseGetInstancesResponse(inputStream io.ReadCloser) InstancesInfo {
    decoder := json.NewDecoder(inputStream)
    var response InstancesInfo
    
    error := decoder.Decode(&response)
    if error != nil {
        fmt.Println("%T\n%s\n%#v\n", error, error, error)
    }
    
    return response
}

func ParseAdbResponse(inputStream io.ReadCloser) AdbResponse {
    decoder := json.NewDecoder(inputStream)
    var response AdbResponse
    
    error := decoder.Decode(&response)
    if error != nil {
        fmt.Println("%T\n%s\n%#v\n", error, error, error)
    }
    
    return response
}