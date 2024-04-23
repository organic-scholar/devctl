package cmd

import (
	"github.com/organic-scholar/devctl/auth"
	"github.com/organic-scholar/devctl/constant"
	"github.com/spf13/cobra"
)

func SshCmd(c *constant.Constant) *cobra.Command {
	return &cobra.Command{
		Use:   "ssh",
		Short: "Manage ssh configurations",
		PreRun: func(cmd *cobra.Command, args []string) {
		},
		Run: func(cmd *cobra.Command, args []string) {
			auth.Authenticate(c)
		},
	}
}

var sshSave = &cobra.Command{
	Use:   "save",
	Short: "Save ssh keys and config",
	Run:   func(cmd *cobra.Command, args []string) {},
}
