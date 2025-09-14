package cmd

import "github.com/jbreckmckye/ff7-tools/ff7"

// CLI definition
// ============================================================================

var Commands struct {
	Code CodeCmds `cmd:"" help:"Work on FFVII code (executable) files"`
}

type CodeCmds struct {
	Extractx ExtractXCmd `cmd:"" help:"Extract MIPS object code from PSX .X files"`
	Patchx   PatchXCmd   `cmd:"" help:"Patch PSX .X file with updated object code"`
}

// ExtractX
// ============================================================================

// ExtractXCmd is the cli definition for extracting object code from .X files like BATTLE.X
type ExtractXCmd struct {
	SOURCE string `arg:"" help:"Path to .X file" type:"existingfile"`
	Dest   string `help:"Where to extract object code (defaults to current dir)"`
}

// Run() is how Kong calls this command
func (cmd *ExtractXCmd) Run() error {
	src := cmd.SOURCE

	// Default dest is SOURCE.obj
	dst := cmd.Dest
	if dst == "" {
		dst = src + ".obj"
	}

	lgr, err := newLogger()
	if err != nil {
		return err
	}

	return ff7.ExtractX(lgr.logMarkdown, src, dst)
}

// PatchX
// ============================================================================

type PatchXCmd struct {
	DEST        string `arg:"" help:"Path to .X file" type:"existingfile"`
	SOURCE      string `arg:"" help:"Path to object code" type:"existingfile"`
	AllowBigger bool   `help:"Disable size checking"`
}
