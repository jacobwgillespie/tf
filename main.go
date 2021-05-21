package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	version, err := versionFromFile()
	if err != nil {
		panic(err)
	}

	if version == "" {
		fmt.Fprintf(os.Stderr, "Unable to find .terraform-version file")
		os.Exit(1)
	}

	execPath, err := binFromVersion(version)
	if err != nil {
		panic(err)
	}

	cmd := exec.Command(execPath, os.Args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err = cmd.Run(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				os.Exit(status.ExitStatus())
			}
		}
		panic(err)
	}
}
