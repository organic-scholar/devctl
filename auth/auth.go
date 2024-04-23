package auth

import "github.com/organic-scholar/devctl/constant"

const oauthUrl = "https://dev-rn476imehrhsh56t.eu.auth0.com/oauth"
const clientId = "qDRn6s4cK0cvjpgb65vDYuABC0Vjqas8"

func Authenticate(c *constant.Constant) {
	_, err := ReadToken(c)
	if err == nil {
		return 
	}
	deviceCode := ReqDeviceCode()
	reqAuthToken(deviceCode.DeviceCode)

}
