package cmd

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/glamour"
)

// mdLogger is a logger that pretty-formats markdown
type mdLogger struct {
	renderer *glamour.TermRenderer
}

func newLogger() (*mdLogger, error) {
	rr, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
	)

	if err != nil {
		return nil, err
	}

	return &mdLogger{
		renderer: rr,
	}, nil
}

func (l *mdLogger) logMarkdown(ss ...string) {
	text := strings.Join(ss, "\n")
	out, err := l.renderer.Render(text)
	if err != nil {
		fmt.Print(text)
	} else {
		fmt.Print(out)
	}
}
