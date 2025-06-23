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

var helloScript []byte
var arg string

// helloCmd represents the hello command
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Run a simple hello world script",

	Run: func(cmd *cobra.Command, args []string) {
		// Write the embedded script to a temporary file
		tmpFile, err := os.CreateTemp("", "gocli-hello-*.sh")
		if err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "Error creating temp file: %v\n", err)
			return
		}
		defer os.Remove(tmpFile.Name())

		_, err = tmpFile.Write(helloScript)
		if err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "Error writing to temp file: %v\n", err)
			return
		}

		err = tmpFile.Chmod(0755)
		if err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "Error setting permissions: %v\n", err)
			return
		}

		tmpFile.Close()

		var execCmd *exec.Cmd

		// Pass the argument to the script if provided
		if arg != "" {
			execCmd = exec.Command("bash", tmpFile.Name(), arg)
		} else {
			execCmd = exec.Command("bash", tmpFile.Name())
		}

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
	runCmd.AddCommand(helloCmd)

	// Add -a / --arg flag
	helloCmd.Flags().StringVarP(&arg, "arg", "a", "", "Argument to pass to hello script")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helloCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
