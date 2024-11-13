package auth

import (
	"fmt"

	"github.com/gistsapp/cli/cmd/auth/actions"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate to Gists.app",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Authenticate to Gists.app")
	},
}

func init() {
	Cmd.AddCommand(actions.LoginCmd, actions.LogoutCmd)
}
