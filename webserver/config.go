package webserver

const baseURL string = "https://api-cloud.genymotion.com/v1"
const authenticationEndpoint string = "/users/signin"
const listInstancesEndpoint string = "/instances"

// return properly built login URL 
func GetLoginURL() string {
    return baseURL + authenticationEndpoint
}

// return properly built intstance URL
func GetListInstancesURL() string {
    return baseURL + listInstancesEndpoint
}