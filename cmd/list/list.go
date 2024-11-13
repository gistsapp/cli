package list

import (
	"fmt"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "list",
	Short: "List all your gists",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Here are all your gists")
	},
}
