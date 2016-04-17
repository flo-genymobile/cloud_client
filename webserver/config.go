package webserver

import "fmt"

const baseURL string = "https://api-cloud.genymotion.com/v1"
const authenticationEndpoint string = "/users/signin"
const listInstancesEndpoint string = "/instances"
const adbEndpoint string = "/instances/%s/adb"

// return properly built login URL 
func GetLoginURL() string {
    return baseURL + authenticationEndpoint
}

// return properly built intstance URL
func GetListInstancesURL() string {
    return baseURL + listInstancesEndpoint
}

func GetAdbURL(instanceID string) string {
    return baseURL + fmt.Sprintf(adbEndpoint, instanceID)
}