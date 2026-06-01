package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var version = "dev"

func getLatestTag() (string, error) {
	tag, err := exec.Command("git", "describe", "--tags", "--abbrev=0").Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(tag)), nil
}

func bumpVersion(tag string, bump string) string {
	version := strings.TrimPrefix(tag, "v")
	parts := strings.Split(version, ".")
	major, _ := strconv.Atoi(parts[0])
	minor, _ := strconv.Atoi(parts[1])
	patch, _ := strconv.Atoi(parts[2])

	switch bump {
	case "major":
		major++
		minor = 0
		patch = 0
	case "minor":
		minor++
		patch = 0
	case "patch":
		patch++
	}
	newTag := fmt.Sprintf("v%d.%d.%d", major, minor, patch)
	return newTag
}

func registerNewTag(newTag string) error {
	cmd := exec.Command("git", "tag", newTag, "--no-sign")
	return cmd.Run()
}

func pushTag(newTag string) error {
	cmd := exec.Command("git", "push", "origin", newTag)
	return cmd.Run()
}

func manualBump() (string, error) {
	var newTag string
	currentTag, err := getLatestTag()
	if err != nil {
		fmt.Println("Error fetching latest tag:", err)
		return "", err
	}

	fmt.Println("Current Tag:", currentTag)
	fmt.Printf("Enter new tag: ")
	fmt.Scanln(&newTag)

	newTag = strings.TrimSpace(newTag)
	newTag = strings.ReplaceAll(newTag, " ", "")
	return newTag, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: git bump <major|minor|patch|manual>")
		return
	}
	argument := os.Args[1]
	if argument != "major" && argument != "minor" && argument != "patch" && argument != "manual" && argument != "version" && argument != "current" {
		fmt.Println("Invalid argument. Use 'major', 'minor', 'patch', 'manual', 'version', or 'current'.")
		return
	}
	latestTag, err := getLatestTag()
	if err != nil && strings.Contains(err.Error(), "128") {
		fmt.Println("No tags found. Creating initial tag v1.0.0.")
		newTag := "v1.0.0"
		_ = registerNewTag(newTag)
		_ = pushTag(newTag)
		os.Exit(0)
	}

	if argument == "major" || argument == "minor" || argument == "patch" {
		newTag := bumpVersion(latestTag, argument)
		err = registerNewTag(newTag)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		err = pushTag(newTag)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("New Tag:", newTag)

	} else if argument == "manual" {
		newTag, err := manualBump()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		err = registerNewTag(newTag)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		err = pushTag(newTag)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("New Tag:", newTag)
	} else if argument == "version" {
		fmt.Println("Current Version:", version)
	} else if argument == "current" {
		fmt.Println("Current Tag:", latestTag)
	}
}
