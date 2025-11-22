[![Go Report Card](https://goreportcard.com/badge/github.com/ansipixels/demo)](https://goreportcard.com/report/github.com/ansipixels/demo)
[![GitHub Release](https://img.shields.io/github/release/ansipixels/demo.svg?style=flat)](https://github.com/ansipixels/demo/releases/)
[![CI Checks](https://github.com/ansipixels/demo/actions/workflows/include.yml/badge.svg)](https://github.com/ansipixels/demo/actions/workflows/include.yml)
[![codecov](https://codecov.io/github/ansipixels/demo/graph/badge.svg?token=CODECOV_TOKEN)](https://codecov.io/github/ansipixels/demo)

# demo

Demo embeds a number of TUI [ansipixels](https://github.com/fortio/terminal/#fortioorgterminalansipixels) program into a single demo binary with a menu

## Install
You can get the binary from [releases](https://github.com/ansipixels/demo/releases)

Or just run
```
CGO_ENABLED=0 go install github.com/ansipixels/demo@latest  # to install (in ~/go/bin typically) or just
CGO_ENABLED=0 go run github.com/ansipixels/demo@latest  # to run without install
```

or
```
docker run -ti ghcr.io/ansipixels/demo
```


## Usage

```
demo help

flags:
```
