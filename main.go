// demo
// Demo embbeds a number of TUI ansipixels program into a single demo binary with a menu

package main

import (
	"os"

	"github.com/ansipixels/demo/cli"
)

func main() {
	os.Exit(cli.Main())
}
