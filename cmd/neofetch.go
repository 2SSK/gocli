/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var neofetchScript []byte

// neofetchCmd represents the neofetch command
var neofetchCmd = &cobra.Command{
	Use:   "neofetch",
	Short: "Display system information using Neofetch",

	Run: func(cmd *cobra.Command, args []string) {
		tmpFile, err := os.CreateTemp("", "gocli-neofetch-*.sh")
		if err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "Error creating temp file: %v\n", err)
			return
		}
		defer os.Remove(tmpFile.Name())

		_, err = tmpFile.Write(neofetchScript)
		if err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "Error writing to temp file: %v\n", err)
			return
		}

		err = tmpFile.Chmod(0755)
		if err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "Error setting permissions: %v\n", err)
			return
		}

		execCmd := exec.Command("bash", tmpFile.Name())

		// Direct the output to CLI
		execCmd.Stdout = cmd.OutOrStdout()
		execCmd.Stderr = cmd.OutOrStderr()

		// Execute the script
		if err := execCmd.Run(); err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "Error executing script: %v\n", err)
			return
		}
	},
}

func init() {
	installCmd.AddCommand(neofetchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// neofetchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// neofetchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
