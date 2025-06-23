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

var arg string

// helloCmd represents the hello command
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		script := "./scripts/hello.sh"

		var execCmd *exec.Cmd

		// Pass the argument to the script if provided
		if arg != "" {
			execCmd = exec.Command("bash", script, arg)
		} else {
			execCmd = exec.Command("bash", script)
		}

		// Direct the outout to CLI
		execCmd.Stdout = cmd.OutOrStdout()
		execCmd.Stderr = cmd.OutOrStderr()

		// Ensure the script is executable
		err := os.Chmod(script, 0755)
		if err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "Error making script executable: %v\n", err)
			return
		}

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
