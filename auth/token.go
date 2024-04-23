package auth

import (
	"encoding/json"
	"time"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	IDToken      string `json:"id_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	CreatedAt    int64  `json:"created_at"`
}

func UnmarshalToken(data []byte) (Token, error) {
	var r Token
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Token) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func (t *Token) IsExpired() bool {
	return t.CreatedAt + t.ExpiresIn < time.Now().Unix() - 60
}
