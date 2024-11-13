package edit

import (
	"fmt"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit one of your gists",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Edit one of your gists")
	},
}
