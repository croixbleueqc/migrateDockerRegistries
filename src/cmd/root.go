// migrateDockerRegistries
// src/cmd/root.go

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"migrateDockerRegistries/connection"
	"migrateDockerRegistries/env"
	"migrateDockerRegistries/helpers"
	"migrateDockerRegistries/img"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "migrateDockerRegistries",
	Short:   "Add a short description here",
	Version: "1.00.00-0 (2024.01.11)",
	Long: `This tools allows you to a software directory structure.
This follows my template and allows you with minimal effort to package your software once built`,
}

var clCmd = &cobra.Command{
	Use:     "changelog",
	Aliases: []string{"cl"},
	Short:   "Shows changelog",
	Run: func(cmd *cobra.Command, args []string) {
		helpers.ChangeLog()
	},
}

var imgLsCmd = &cobra.Command{
	Use:     "ls",
	Aliases: []string{"compare"},
	Short:   "Compare lists",
	Run: func(cmd *cobra.Command, args []string) {
		if err := img.CompareImagesLists(); err != nil {
			fmt.Println("Error:", err)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.DisableAutoGenTag = true
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(clCmd, imgLsCmd)

	rootCmd.PersistentFlags().StringVarP(&env.EnvConfigFile, "env", "e", "defaultEnv.json", "Default environment configuration file; this is a per-user setting.")
	rootCmd.PersistentFlags().StringVarP(&connection.ConnectURI, "host", "H", "unix:///var/run/docker.sock", "Remote host:port to connect to")
}
