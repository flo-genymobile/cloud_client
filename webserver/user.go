package webserver

type User struct {
    Email string `json:"email"`
    Password string `json:"password"`
}

type SigninData struct {
    Message string `json:"message"`
    Token string `json:"token"`
}

type SigninResponse struct {
    Data SigninData `json:"data"`
    Version float64 `json:"version"`
}