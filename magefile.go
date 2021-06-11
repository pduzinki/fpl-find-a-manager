//+build mage

package main

import "github.com/magefile/mage/sh"

func Build() error {
	// go mod download
	if err := sh.Run("go", "mod", "download"); err != nil {
		return err
	}

	// GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o app ./cmd/cli/
	env := map[string]string{
		"GOOS":   "linux",
		"GOARCH": "amd64",
	}
	return sh.RunWith(env, "go", "build", "-ldflags="+"-w -s", "-o", "app", "./cmd/cli/")
}