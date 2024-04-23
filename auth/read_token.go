package auth

import (
	"os"

	"github.com/organic-scholar/devctl/constant"
)

func ReadToken(c *constant.Constant) (token Token, err error) {
	bytes, err := os.ReadFile(c.TokenJsonFile)
	if err != nil {
		return
	}
	token, err = UnmarshalToken(bytes)
	if err != nil || token.IsExpired() {
		os.Remove(c.TokenFile)
		os.Remove(c.TokenJsonFile)
	}
	return token, err
}
