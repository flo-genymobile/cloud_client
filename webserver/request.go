package webserver

import (
    "fmt"
    "bytes"
    "net/http"
    "encoding/json"
    "io"
    "mime/multipart"
    "os"
    "path/filepath"
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

func PreparePost(url string, token string) *http.Request {
    request, error := http.NewRequest("POST", url, nil)
    if error != nil {
        fmt.Println("can't create request: ", error)
    }
    
    request.Header.Set("Authorization", "Bearer " + token)
    return request
}

func PrepareAdbPost(url string, token string, commandInfo AdbCommandInfo) *http.Request {
    //Prepare multi part writer
    body := &bytes.Buffer{}
    writer := multipart.NewWriter(body)
    writer.WriteField("cmd", commandInfo.Command)
    writer.WriteField("args", commandInfo.Arguments)
    
    if commandInfo.FilePath != "" {
        //Open file
        file, error := os.Open(commandInfo.FilePath)
        if error != nil {
            fmt.Println("can't create request: ", error)
            return nil
        }
        defer file.Close()
        
        //add raw file data
        part, error := writer.CreateFormFile("file", filepath.Base(commandInfo.FilePath))
        if error != nil {
            fmt.Println("can't create request: ", error)
            return nil
        }
        io.Copy(part, file)
    }
    
    error := writer.Close()
        if error != nil {
            fmt.Println("can't create request: ", error)
            return nil
        }
  
    request, error := http.NewRequest("POST", url, body)
    if error != nil {
        fmt.Println("can't create request: ", error)
    }
    
    request.Header.Set("Content-Type", writer.FormDataContentType())
    request.Header.Set("Authorization", "Bearer " + token)
    
    return request
}

func DoRequest(request *http.Request) io.ReadCloser {
    httpClient := &http.Client{}
    fmt.Println("Do request: " + request.Method + ", " + request.Host + ", " + request.URL.Path)
    httpresponse, error := httpClient.Do(request)
    if error != nil {
        fmt.Println("Request failed: ", error)
    }
    
    //TODO: where does this goes...?
    //defer httpresponse.Body.Close()
    
    //TODO: handle errors
    fmt.Println("Request Status Code: ", httpresponse.StatusCode)
    fmt.Println()   
    
    return httpresponse.Body    
}