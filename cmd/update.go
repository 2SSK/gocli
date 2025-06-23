/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update gocli to the latest version",
	Run: func(cmd *cobra.Command, args []string) {
		owner := "2SSK"
		repo := "gocli"
		apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", owner, repo)

		fmt.Println("Checking for latest version...")

		resp, err := http.Get(apiURL)
		if err != nil {
			fmt.Printf("Failed to fetch release info: %v\n", err)
			return
		}
		defer resp.Body.Close()

		var release struct {
			TagName string `json:"tag_name"`
			Assets  []struct {
				Name               string `json:"name"`
				BrowserDownloadURL string `json:"browser_download_url"`
			} `json:"assets"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
			fmt.Printf("Error parsing release data: %v\n", err)
			return
		}

		osName := runtime.GOOS
		arch := runtime.GOARCH

		binaryName := fmt.Sprintf("gocli-%s-%s", osName, arch)

		var downloadURL string
		for _, asset := range release.Assets {
			if asset.Name == binaryName {
				downloadURL = asset.BrowserDownloadURL
				break
			}
		}

		if downloadURL == "" {
			fmt.Printf("No binary found for %s-%s\n", osName, arch)
			return
		}

		fmt.Printf("Downloading update from %s...\n", downloadURL)

		resp, err = http.Get(downloadURL)
		if err != nil {
			fmt.Printf("Failed to download binary: %v\n", err)
			return
		}
		defer resp.Body.Close()

		// Save to temporary file
		tmpFile := "gocli_tmp"
		out, err := os.Create(tmpFile)
		if err != nil {
			fmt.Printf("Failed to create file: %v\n", err)
			return
		}
		defer out.Close()

		_, err = io.Copy(out, resp.Body)
		if err != nil {
			fmt.Printf("Error saving file: %v\n", err)
			return
		}

		err = os.Chmod(tmpFile, 0755)
		if err != nil {
			fmt.Printf("Failed to set executable permission: %v\n", err)
			return
		}

		// Get the path to the currently running executable
		currentBinary, err := os.Executable()
		if err != nil {
			fmt.Printf("Failed to find current binary: %v\n", err)
			return
		}

		// Replace current binary
		err = os.Rename(tmpFile, currentBinary)
		if err != nil {
			fmt.Printf("Failed to update binary: %v\n", err)
			return
		}

		fmt.Printf("gocli updated successfully to version %s!\n", release.TagName)

		// Optional: Restart the CLI if you want
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
