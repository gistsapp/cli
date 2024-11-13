package actions

import (
	"fmt"

	"github.com/spf13/cobra"
)

var LogoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout to Gists.app",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Logout to Gists.app")
	},
}
