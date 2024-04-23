package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/organic-scholar/devctl/cmd"
	"github.com/organic-scholar/devctl/constant"
	"github.com/organic-scholar/devctl/di"
)

//go:embed pages/*.html
var pagesFs embed.FS

func main() {
	cont := di.NewContainer()
	cont.Provide(constant.ProvideConstant)
	c := di.Get[*constant.Constant](cont)
	rootCmd := cmd.RootCmd()
	rootCmd.AddCommand(cmd.SshCmd(c))
	rootCmd.AddCommand(cmd.ServerCmd(c, pagesFs))
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
