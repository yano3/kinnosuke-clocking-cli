# kinnosuke-clocking-cli

[![GitHub release](https://img.shields.io/github/release/yano3/kinnosuke-clocking-cli.svg)](https://github.com/yano3/kinnosuke-clocking-cli/releases)
[![CircleCI](https://circleci.com/gh/yano3/kinnosuke-clocking-cli.svg?style=shield)](https://circleci.com/gh/yano3/kinnosuke-clocking-cli)

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

### macOS

If you use [Homebrew](https://brew.sh):

```
brew tap yano3/tap
brew install kinnosuke-clocking-cli
```

### Other platforms

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
