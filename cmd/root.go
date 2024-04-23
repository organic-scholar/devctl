package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func RootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "devctl",
		Short: "devctl is tool for doing repetitive things",
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			// println("hello")
			// res, err := auth.ReqDeviceCode()
			// if err != nil {
			// 	log.Fatal(err.Error())
			// }
			// tokenRes := auth.ReqAuthToken(res.DeviceCode)
			// println(tokenRes.AccessToken)
		},
	}

}
