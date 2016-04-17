package webserver

import (
    "fmt"
    "bytes"
    "net/http"
    "encoding/json"
    "io"
)

func PrepareLogin(url string, user User) *http.Request {
    payload, error := json.Marshal(user)
    if error != nil {
        fmt.Println("can't marhsall: ", error)
    }
    
    request, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
    if err != nil {
        fmt.Println("can't create request: ", err)
    }
    request.Header.Set("Content-Type", "application/json")
    
    return request
}

func PrepareGet(url string, token string) *http.Request {
    request, error := http.NewRequest("GET", url, nil)
    if error != nil {
        fmt.Println("can't create request: ", error)
    }
    
    request.Header.Set("Authorization", "Bearer " + token)
    return request
}

func DoRequest(request *http.Request) io.ReadCloser {
    httpClient := &http.Client{}
    httpresponse, error := httpClient.Do(request)
    if error != nil {
        fmt.Println("Request failed: ", error)
    }
    
    //TODO: where does this goes...?
    //defer httpresponse.Body.Close()
    
    //TODO: handle errors
    fmt.Println("Request Status Code: ", httpresponse.StatusCode)
    
    return httpresponse.Body    
}