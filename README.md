## Golang client for GenymotionCLOUD

This client requires a valid username and password with an acitvated license. </br>
All calls to the https://api-cloud.genymotion.com/v1 require a valid token obtained after a succesfull login on the cloud. 

The GenymotionCloud documentation can be found here: https://doc-api-cloud-staging.genymotion.com/index.html

### How to deploy

This client is written in go. In order to execute it you need a proper go environnement setup.  </br>
1. You can download go here: https://golang.org/dl/ </br>
2. Setup you go workspace by defining the GOPATH environnement variable. </br>
 ```$ export GOPATH=<path/to/workspace> ```  </br>
3. Clone this repository inside your go workspace. </br>
4. Build and install the app </br>
``` $ go install cloud_client ``` </br>
5. If all goes well a /bin directory with the binary should have appeared in your go workspace. </br>
6. run it! </br>

### What has been implemented so far

- Signin works and token is extracted </br>
https://doc-api-cloud-staging.genymotion.com/endpoints/api-auth.html#post-v1-users-signin

- Listing of currently running VMs is implemented </br>
https://doc-api-cloud-staging.genymotion.com/endpoints/api-vmmanage.html#get-v1-instances

- Adb push and Adb install can be performed on a running VMs </br>
https://doc-api-cloud-staging.genymotion.com/endpoints/api-vmmanage.html#post-v1-instances--instance_id--adb

