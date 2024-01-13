// nxrmuploader
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/cmd/envCommands.go
// Original timestamp: 2023/12/31 14:44

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"migrateDockerRegistries/env"
	"os"
)

var envCmd = &cobra.Command{
	Use:     "env",
	Aliases: []string{"environment"},
	Short:   "Environment subcommands",
	Long:    `You need to provide the subcommands: ls, create, rm, explain.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You need to provide one of the following subcommands: ls, create, rm or explain")
	},
}

var envLsCmd = &cobra.Command{
	Use:     "ls",
	Aliases: []string{"list"},
	Short:   "List environment files",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if len(args) > 0 {
			err = env.ListEnvironments(args[0])
		} else {
			err = env.ListEnvironments("")
		}
		if err != nil {
			fmt.Println(err)
		}
	},
}

var envRmCmd = &cobra.Command{
	Use:     "rm",
	Aliases: []string{"del", "remove"},
	Short:   "Delete environment file(s)",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if len(args) == 0 {
			fmt.Println("You need to provide at least one filename")
			os.Exit(0)

		} else {
			if err = env.RemoveEnvFile(args); err != nil {
				fmt.Println(err)
			}
		}
	},
}

var envAddCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"create"},
	Short:   "Create an environment file",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var reponame string
		if len(args) == 0 {
			reponame = "defaultEnv.json"
		} else {
			reponame = args[0]
		}
		if err = env.AddEnvFile(reponame); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(envCmd)
	envCmd.AddCommand(envLsCmd, envRmCmd, envAddCmd)
}
