package auth

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"slices"
	"strings"
	"time"
)

func ReqAuthToken(deviceCode string) Token {
	for {
		res, err := reqAuthToken(deviceCode)
		if err == nil {
			return res
		}
		if slices.Contains([]string{"authorization_pending", "slow_down"}, err.Error()) {
			time.Sleep(5 * time.Second)
			continue
		}
		log.Fatal(err)
	}
}

func reqAuthToken(deviceCode string) (token Token, err error) {
	params := url.Values{}
	params.Set("client_id", clientId)
	params.Set("device_code", deviceCode)
	params.Set("grant_type", "urn:ietf:params:oauth:grant-type:device_code")
	req, _ := http.NewRequest("POST", oauthUrl+"/oauth/token", strings.NewReader(params.Encode()))
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	if slices.Contains([]int{403, 429}, res.StatusCode) {
		errorResponse, err := UnmarshalErrorResponse(body)
		fmt.Println(errorResponse.ErrorDescription)
		if err != nil {
			log.Fatal(err)
		}
		return token, errorResponse
	}
	if res.StatusCode != 200 {
		log.Print("Invalid Response")
		log.Fatal(res)
	}
	token, err = UnmarshalToken(body)
	token.CreatedAt = time.Now().Unix()
	return token, err
}
