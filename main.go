package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func getChangedFiles(baseCommit, headCommit string) ([]string, error) {
	cmd := exec.Command("git", "diff", "--name-only", fmt.Sprintf("%s..%s", baseCommit, headCommit))
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error running git diff: %v", err)
	}
	files := strings.Split(strings.TrimSpace(string(output)), "\n")
	return files, nil
}

func findChangedApps(changedFiles []string, appsFolder string) map[string]struct{} {
	changedApps := make(map[string]struct{})

	for _, filePath := range changedFiles {
		normalizedPath := filepath.ToSlash(filePath)
		if strings.HasPrefix(normalizedPath, appsFolder+"/") {
			relativePath := strings.TrimPrefix(normalizedPath, appsFolder+"/")
			appName := strings.Split(relativePath, "/")[0]
			changedApps[appName] = struct{}{}
		}
	}

	return changedApps
}

func main() {
	baseCommit := flag.String("base", "HEAD^", "Base commit to compare (default: HEAD^)")
	headCommit := flag.String("head", "HEAD", "Head commit to compare (default: HEAD)")
	appsFolder := flag.String("projects", "apps", "Folder containing the projects (default: apps)")

	flag.Parse()

	if _, err := os.Stat(*appsFolder); os.IsNotExist(err) {
		fmt.Printf("Error: '%s' folder not found. Make sure you're in the root of an Nx monorepo.\n", *appsFolder)
		os.Exit(1)
	}

	changedFiles, err := getChangedFiles(*baseCommit, *headCommit)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	if len(changedFiles) == 0 {
		fmt.Println("No changes detected.")
		return
	}

	changedApps := findChangedApps(changedFiles, *appsFolder)
	if len(changedApps) > 0 {
		for app := range changedApps {
			fmt.Printf("%s", app)
		}
	} else {
		fmt.Printf("No apps in the '%s' folder have changed.\n", *appsFolder)
		os.Exit(1)
	}
}
