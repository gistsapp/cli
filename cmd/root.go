package cmd

import (
	"fmt"
	"os"

	"github.com/gistsapp/cli/cmd/create"
	"github.com/gistsapp/cli/cmd/auth"
	"github.com/gistsapp/cli/cmd/edit"
	"github.com/gistsapp/cli/cmd/exec"
	"github.com/gistsapp/cli/cmd/get"
	"github.com/gistsapp/cli/cmd/list"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gists",
	Short: "gists CLI",
	Long:  "gists is a CLI to interact with Gists.app",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Bienvenue dans mon CLI ! Utilisez --help pour voir les commandes disponibles.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(auth.Cmd, get.Cmd, list.Cmd, edit.Cmd, exec.Cmd, create.Cmd)
}
