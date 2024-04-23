package main

import (
	"fmt"
	"os"

	"github.com/organic-scholar/devctl/cmd"
	"github.com/organic-scholar/devctl/constant"
	"github.com/organic-scholar/devctl/di"
)

func main() {
	cont := di.NewContainer()
	cont.Provide(constant.ProvideConstant)
	c := di.Get[*constant.Constant](cont)
	rootCmd := cmd.RootCmd()
	sshCmd := cmd.SsCmd(c)
	sshCmd.AddCommand(cmd.SshSaveCommand(c))
	rootCmd.AddCommand(sshCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
