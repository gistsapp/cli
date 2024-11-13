package actions

import (
	"fmt"

	"github.com/spf13/cobra"
)

var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to Gists.app",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Login to Gists.app")
	},
}
