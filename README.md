# kinnosuke-clocking-cli
Clocking in/out cli for Kinnosuke.

## Usage

Clocking in

```
kinnosuke-clocking-cli
```

Clocking out

```
kinnosuke-clocking-cli -out
```

## Installation

Download binary from [releases page](https://github.com/yano3/kinnosuke-clocking-cli/releases) or use `go get` command.

```console
$ go get -u github.com/yano3/kinnosuke-clocking-cli
```

## Configuration

Set environment variables bellow.

```
export KINNOSUKE_COMPANYCD=<Put your customer id>
export KINNOSUKE_LOGINCD=<Put your login id (typically your employee number)>
export KINNOSUKE_PASSWORD=<Put your password>
```
