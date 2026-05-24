package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

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

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: git bump <major|minor|patch>")
		return
	}
	argument := os.Args[1]
	if argument != "major" && argument != "minor" && argument != "patch" {
		fmt.Println("Invalid argument. Use 'major', 'minor', or 'patch'.")
		return
	}
	latestTag, err := getLatestTag()
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("No tags found. Please create an initial tag before bumping.")
		return
	}
	newTag := bumpVersion(latestTag, argument)
	err = registerNewTag(newTag)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("New Tag:", newTag)
}
