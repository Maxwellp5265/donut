package main

import (
	"flag"
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/erik-adelbert/donut/donut"
	"github.com/erik-adelbert/donut/pkg/epilepsy"
	"golang.org/x/term"
)

func main() {
	noWarning := flag.Bool("no-warning", false, "Skip the epilepsy warning screen")
	flag.Parse()

	if !*noWarning {
		if ok := epilepsy.Warn(); !ok {
			return
		}
	}

	w, h, err := term.GetSize(int(os.Stdin.Fd()))

	if err != nil {
		fatal("Could not get terminal size:", err)
	}

	h = max(1, h) // ensure the dimensions are strictly positive
	w = max(1, w)
	p := tea.NewProgram(donut.NewModel(h, w))

	if _, err := p.Run(); err != nil {
		fatal("Error running program:", err)
	}
}

func fatal(a ...any) {
	fmt.Fprintln(os.Stderr, a...)
	os.Exit(1)
}
