package create

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "create",
	Short: "Create one of your gists",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Create one of your gists")
	},
}
