package main

import (
	"bytes"
	"os/exec"
	"strings"
)

func parseID(url string) string {
	// Guess at id by parsing last segment of the first remote URL
	// TODO look into using first commit ID or something else more conclusive
	url = strings.TrimSuffix(url, ".git")
	parts := strings.Split(url, "/")
	return parts[len(parts)-1]
}

func parseRemotes(cfg string) []string {
	remotes := []string{}
	for _, row := range strings.Split(cfg, "\n") {
		if !strings.HasPrefix(row, "remote") {
			continue
		}
		parts := strings.Split(row, "=")
		if strings.HasSuffix(parts[0], "url") {
			remotes = append(remotes, parts[1])
		}
	}
	return remotes
}

func git(path string, args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
	)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Dir = path
	err := cmd.Run()
	return stdout.String() + stderr.String(), err
}
