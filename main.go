package main

import (
	"github.com/alecthomas/kong"

	"github.com/jbreckmckye/ff7-tools/cmd"
)

func main() {
	ctx := kong.Parse(&cmd.Commands)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
