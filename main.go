package main

import (
	"github.com/alecthomas/kong"
)

type Globals struct {
	Debug bool `help:"Log debug statements"`
}

var commands struct {
	Globals
	Code CodeCmds `cmd:"" help:"Work on FFVII code (executable) files"`
}

type CodeCmds struct {
	ExtractX ExtractXCmd `cmd:"" help:"Extract MIPS object code from PSX .X files"`
	PatchX PatchXCmd `cmd:"" help:"Patch PSX .X file with updated object code"`
}

type ExtractXCmd struct {
  SOURCE string `arg:"" help:"Path to .X file" type:"existingfile"`
  Dest string `help:"Where to extract object code (defaults to current dir)" type:"existingFile"`
}

type PatchXCmd struct {
	DEST string `arg:"" help:"Path to .X file" type:"existingfile"`
	SOURCE string `arg:"" help:"Path to object code" type:"existingfile"`
	AllowBigger bool `help:"Disable size checking"`
}

func main() {
  ctx := kong.Parse(&commands)
  globals := &Globals{
		Debug: commands.Debug,
	}
	err := ctx.Run(globals)
	ctx.FatalIfErrorf(err)
}
