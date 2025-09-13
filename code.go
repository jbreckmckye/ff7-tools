package main

import (
	"errors"
	"fmt"
)

func (cmd *ExtractXCmd) Run(g *Globals) error {
	fmt.Printf("inspect stuff dest %v src %v dbg %v\n", cmd.Dest, cmd.SOURCE, g.Debug)
	return nil
}

func (cmd *PatchXCmd) Run() error {
	return fmt.Errorf("this is a test error: %w", errors.New("i'm a bad man"))
}
