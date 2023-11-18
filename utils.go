package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/go-version"
	install "github.com/hashicorp/hc-install"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/hc-install/src"
)

const versionFileName = ".terraform-version"

func binFromVersion(versionString string) (string, error) {
	ctx := context.Background()

	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	execPath := filepath.Join(homedir, ".tf", fmt.Sprintf("terraform-%s", versionString))

	_, err = os.Stat(execPath)
	if errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(filepath.Join(homedir, ".tf"), 0755)
		if err != nil {
			return "", err
		}

		i := install.NewInstaller()
		defer i.Remove(ctx)

		downloadPath, err := i.Install(ctx, []src.Installable{
			&releases.ExactVersion{
				Product: product.Terraform,
				Version: version.Must(version.NewVersion(versionString)),
			},
		})
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
		files, err := os.ReadDir(dir)
		if err != nil {
			return "", err
		}

		for _, f := range files {
			if f.Name() == versionFileName {
				data, err := os.ReadFile(filepath.Join(dir, versionFileName))
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
