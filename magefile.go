//+build mage

package main

import (
	"errors"
	"fmt"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var ErrUnknownTarget error = errors.New("Unknown target!")

func Clear() error {
	return sh.Run("rm", "app", "-f")
}

func Build(what string) error {
	if what != "cli" && what != "server" {
		fmt.Printf("No such target as '%s'. Use 'cli' or 'server' instead.\n", what)
		return ErrUnknownTarget
	}

	mg.Deps(Clear)

	// go mod download
	if err := sh.Run("go", "mod", "download"); err != nil {
		return err
	}

	// GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o app ./cmd/cli/
	env := map[string]string{
		"GOOS":   "linux",
		"GOARCH": "amd64",
	}
	return sh.RunWith(env, "go", "build", "-ldflags="+"-w -s", "-o", "app",
		fmt.Sprintf("./cmd/%s/", what))
}
