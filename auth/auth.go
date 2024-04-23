package auth

import (
	"fmt"

	"github.com/organic-scholar/devctl/constant"
)

const oauthUrl = "https://dev-rn476imehrhsh56t.eu.auth0.com/oauth"
const clientId = "8uwhNiVKrkXKk20K8ZuFGuRamkRAHAyD"

func Authenticate(c *constant.Constant) {
	_, err := ReadToken(c)
	if err == nil {
		return
	}
	deviceCode := ReqDeviceCode()
	printDeviceCodeMessage(deviceCode)
	ReqAuthToken(deviceCode.DeviceCode)

}

func printDeviceCodeMessage(deviceCode DeviceCode) {
	fmt.Println("Attempting to automatically open the SSO authorization page in your default browser.")
	fmt.Println("If the browser does not open or you wish to use a different device to authorize this request, open the following URL:")
	fmt.Println()
	fmt.Println(deviceCode.VerificationURIComplete)
}
