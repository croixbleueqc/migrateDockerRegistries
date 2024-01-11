// migrateDockerRegistries
// src/cmd/root.go

package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "migrateDockerRegistries",
	Short:   "Add a short description here",
	Version: "1.00.00-0 (2024.xx.yy)",
	Long: `This tools allows you to a software directory structure.
This follows my template and allows you with minimal effort to package your software once built`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
