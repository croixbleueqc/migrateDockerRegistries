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
	Short:   "Docker registries migration tool",
	Version: "1.10.00-0 (2024.01.17)",
	Long: `This tool lists all images+tags from a source docker registry, then does the same on a destination registry.
Once both lists compiled, a third list is generated with all the images+tags missing in the second registry.
Then, optionally, you can have a shell script that retags all those missing images+tags, and (another option) create the push command.`,
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
	Use:     "compare",
	Aliases: []string{"ls"},
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
	imgLsCmd.PersistentFlags().BoolVarP(&img.Retag, "retag", "r", false, "Create a command (shell) file retagging all missing images to the new repo name")
	imgLsCmd.PersistentFlags().BoolVarP(&img.Push, "push", "p", false, "Add a docker push command to the above shell file")
	imgLsCmd.PersistentFlags().BoolVarP(&img.Delete, "delete", "d", false, "Delete the images after retag")
	imgLsCmd.PersistentFlags().BoolVarP(&img.LatestOnly, "latest", "l", false, "Only fetches the 'latest' tag from the source registry")
}
