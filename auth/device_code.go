package auth

import "encoding/json"

type DeviceCode struct {
	DeviceCode              string `json:"device_code"`
	UserCode                string `json:"user_code"`
	VerificationURI         string `json:"verification_uri"`
	VerificationURIComplete string `json:"verification_uri_complete"`
	ExpiresIn               int64  `json:"expires_in"`
	Interval                int64  `json:"interval"`
}

func UnmarshalDeviceCode(data []byte) (DeviceCode, error) {
	var r DeviceCode
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DeviceCode) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
