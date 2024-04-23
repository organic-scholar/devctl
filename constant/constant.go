package constant

import (
	"os"
	"path"
)

type Constant struct {
	UserHomeDir   string
	SshDir        string
	TokenFile     string
	TokenJsonFile string
}

func ProvideConstant() *Constant {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return &Constant{
		UserHomeDir:   homeDir,
		SshDir:        path.Join(homeDir, ".ssh"),
		TokenFile:     path.Join(homeDir, ".devctl", "token"),
		TokenJsonFile: path.Join(homeDir, ".devctl", "token.json"),
	}
}
