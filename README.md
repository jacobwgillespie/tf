# tf

A small `terraform` wrapper which automatically downloads and executes specific Terraform versions.

## Installation

Either install with Homebrew, or download from the [project releases](https://github.com/jacobwgillespie/tf/releases):

```shell
$ brew install jacobwgillespie/tap/tf
```

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

### Default Terraform Version

`tf` will print an error if it is unable to find a `.terraform-version` file. To set the default Terraform version, create a `.terraform-version` file in your home directory, then all directories underneath will default to your specified version.

## License

MIT License, see `LICENSE`
