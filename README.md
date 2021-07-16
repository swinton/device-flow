# `device-flow`

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
./bin/device-flow 178c6fc778ccc68e1d6a -S repo -S read:org
```
