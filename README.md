# tf

A small `terraform` wrapper which automatically downloads and executes specific Terraform versions.

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

## License

MIT License, see `LICENSE`
