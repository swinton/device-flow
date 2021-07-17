# :v: `device-flow` ![](https://github.com/swinton/device-flow/workflows/Build/badge.svg)
> Generate access tokens from the command line with the OAuth device authorization flow

## Build
Run `make`.

```shell
make
```

This will drop a binary file, `device-flow`, in the `bin` folder.

## Usage
```shell
./bin/device-flow --help
```

## Example
To generate an access token with `repo` and `read:org` scopes for the OAuth application with client id `178c6fc778ccc68e1d6a`:

```shell
./bin/device-flow auth 178c6fc778ccc68e1d6a -S repo -S read:org
```
