package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/terraform-exec/tfinstall"
)

const versionFileName = ".terraform-version"

func binFromVersion(version string) (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	execPath := filepath.Join(homedir, ".tf", fmt.Sprintf("terraform-%s", version))

	_, err = os.Stat(execPath)
	if errors.Is(err, os.ErrNotExist) {
		if err = os.Mkdir(filepath.Join(homedir, ".tf"), 0755); err != nil && !errors.Is(err, os.ErrExist) {
			return "", err
		}

		tmpDir, err := ioutil.TempDir("", "tfinstall")
		if err != nil {
			return "", err
		}
		defer os.RemoveAll(tmpDir)

		downloadPath, err := tfinstall.Find(context.Background(), tfinstall.ExactVersion(version, tmpDir))
		if err != nil {
			return "", err
		}

		err = os.Rename(downloadPath, execPath)
		if err != nil {
			return "", err
		}
	}

	return execPath, nil
}

func versionFromFile() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			return "", err
		}

		for _, f := range files {
			if f.Name() == versionFileName {
				data, err := ioutil.ReadFile(filepath.Join(dir, versionFileName))
				if err != nil {
					return "", err
				}
				return strings.TrimSpace(string(data)), nil
			}
		}

		parent := filepath.Dir(dir)

		if parent == dir {
			return "", nil
		}

		dir = parent
	}
}
