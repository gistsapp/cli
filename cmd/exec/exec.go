package exec

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "exec",
	Short: "Exec one of your gists",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Exec one of your gists")
	},
}
