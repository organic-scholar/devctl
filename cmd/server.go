package cmd

import (
	"fmt"
	"io/fs"
	"net/http"

	"github.com/organic-scholar/devctl/constant"
	"github.com/organic-scholar/devctl/pages"
	"github.com/spf13/cobra"
)

func ServerCmd(c *constant.Constant, f fs.FS) *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "Run http server",
		PreRun: func(cmd *cobra.Command, args []string) {
		},
		Run: func(cmd *cobra.Command, args []string) {
			r := pages.Router(f)
			port := 8000
			fmt.Printf("Starting server at port %d", port)
			if err := http.ListenAndServe(fmt.Sprintf(":%d", port), r); err != nil {
				fmt.Println("Error starting server:", err)
			}
		},
	}
}
