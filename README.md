# tf

A small `terraform` wrapper which automatically downloads and executes specific Terraform versions.

## Installation

Either download from the [project releases](https://github.com/jacobwgillespie/tf/releases), or install with Homebrew, choosing either the `tf` or `tf-shim` formulae:

- **tf** (installs as `tf`)

  ```shell
  $ brew install jacobwgillespie/tap/tf
  ```

- **tf-shim** (installs as `terraform`)

  ```shell
  $ brew install jacobwgillespie/tap/tf-shim
  ```

**NOTE:** it's recommended to install the `tf-shim` or otherwise alias `terraform` to `tf`, as that will enable automated tools like the Terraform VS Code extension to make use of the wrapper.

## Usage

First, create a file named `.terraform-version` containing your desired Terraform version:

```
0.15.0
```

Next, run `tf` instead of `terraform`:

```shell
$ tf [...]
```

`tf` will automatically download and cache the specified Terraform version to `~/.tf`.

**NOTE:** `tf` itself is silent, it will only print output from the proxied `terraform` binary unless there is an error. The first time you run `tf`, it will take time for it to download the Terraform CLI binary to `~/.tf`.

### Default Terraform Version

`tf` will print an error if it is unable to find a `.terraform-version` file. To set the default Terraform version, create a `.terraform-version` file in your home directory, then all directories underneath will default to your specified version.

### How It Works

`tf` is a simple wrapper around Terraform that executes the version of the Terraform CLI specified in the nearest `.terraform-version` file. Specifically, it:

1. Looks for the nearest `.terraform-version` file, starting in the current working directory and scanning up the file tree.
2. Downloads the specified version to `~/.tf/terraform-VERSION` if not already downloaded using [tfinstall](github.com/hashicorp/terraform-exec)
3. Execs that binary with all passed environment and CLI arguments.

## License

MIT License, see `LICENSE`
