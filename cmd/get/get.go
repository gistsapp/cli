package get

import (
	"fmt"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "get",
	Short: "Get one of your gists",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Here is one of your gists")
	},
}
