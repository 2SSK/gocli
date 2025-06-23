package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

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
		if osName == "windows" {
			binaryName += ".exe"
		}

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

		currentBinary, err := os.Executable()
		if err != nil {
			fmt.Printf("Failed to find current binary: %v\n", err)
			return
		}

		// Prepare bootstrap script
		bootstrapScript := fmt.Sprintf(`#!/bin/bash
sleep 1
mv %s %s
echo "Update completed successfully!"
`, tmpFile, currentBinary)

		scriptPath := "gocli_update.sh"
		scriptFile, err := os.Create(scriptPath)
		if err != nil {
			fmt.Printf("Failed to create update script: %v\n", err)
			return
		}
		defer scriptFile.Close()

		_, err = scriptFile.WriteString(bootstrapScript)
		if err != nil {
			fmt.Printf("Failed to write update script: %v\n", err)
			return
		}

		err = os.Chmod(scriptPath, 0755)
		if err != nil {
			fmt.Printf("Failed to set script as executable: %v\n", err)
			return
		}

		fmt.Println("Launching update script...")

		cmdExec := exec.Command("bash", scriptPath)
		cmdExec.Stdout = os.Stdout
		cmdExec.Stderr = os.Stderr

		if err := cmdExec.Start(); err != nil {
			fmt.Printf("Failed to start update script: %v\n", err)
			return
		}

		fmt.Println("Update script running. Exiting current process...")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
