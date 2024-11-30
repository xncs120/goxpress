package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	repoURL        = "https://github.com/xncs120/goxpress.git"
	oldProjectName = "github.com/xncs120/goxpress"
	removeList     = []string{
		"install.go",
		"LICENSE",
		"README.md",
	}
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run github.com/xncs120/goxpress@main project-name")
		os.Exit(1)
	}
	newProjectName := os.Args[1]

	fmt.Printf("=> Scan for %s folder", newProjectName)
	_, err := os.Stat(newProjectName)
	if !os.IsNotExist(err) {
		fmt.Printf("\t%s folder already present", newProjectName)
		os.Exit(1)
	}

	fmt.Println("=> Cloning repository")
	if err := gitClone(repoURL, newProjectName); err != nil {
		fmt.Printf("\tError cloning repository: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("=> Scaffolding files and folders")
	if err := scaffoldingFilesAndFolders(newProjectName, removeList); err != nil {
		fmt.Printf("\tError when scaffolding up files and folders: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("=> Project [%s] successfully created\n", newProjectName)
}

func gitClone(repoURL, folder string) error {
	cmd := exec.Command("git", "clone", repoURL, folder)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func scaffoldingFilesAndFolders(basePath string, removeList []string) error {
	return filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && filepath.Base(path) == "assets" {
			return filepath.SkipDir
		}

		relPath, _ := filepath.Rel(basePath, path)
		for _, exclude := range removeList {
			if relPath == exclude || filepath.Base(relPath) == exclude {
				if info.IsDir() {
					return os.RemoveAll(path)
				}
				return os.Remove(path)
			}
		}

		if !info.IsDir() {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			contentStr := string(content)
			if strings.Contains(contentStr, oldProjectName) {
				replacedContent := strings.ReplaceAll(contentStr, oldProjectName, basePath)
				err := os.WriteFile(path, []byte(replacedContent), info.Mode())
				if err != nil {
					return err
				}
			}
			if strings.Contains(contentStr, "APP_NAME=goxpress") {
				replacedContent := strings.ReplaceAll(contentStr, "APP_NAME=goxpress", "APP_NAME="+basePath)
				err := os.WriteFile(path, []byte(replacedContent), info.Mode())
				if err != nil {
					return err
				}
			}
		}

		if filepath.Base(path) == ".env.example" {
			newPath := filepath.Join(filepath.Dir(path), ".env")
			return os.Rename(path, newPath)
		}

		return nil
	})
}
