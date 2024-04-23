package auth

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func ReqDeviceCode() (DeviceCode) {
	params := url.Values{}
	params.Set("client_id", clientId)

	req, err := http.NewRequest("POST", oauthUrl+"/device/code", strings.NewReader(params.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	println(string(body))
	deviceCode, err := UnmarshalDeviceCode(body)
	if err != nil {
		log.Fatal(err)
	}
	return deviceCode
}
