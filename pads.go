package main

import (
	"os"
)

type pads struct {
	Head string
	Tail string
}

type padCollection map[string]pads

var padList = padCollection{
	"default": {"\033[5i", "\033[4i"},
	"tmux":    {"\033Ptmux;\033\033[5i", "\033\033[4i\033\\"},
}

func (p padCollection) Current() pads {
	if len(os.Getenv("TMUX")) > 0 {
		return p["tmux"]
	}
	return p["default"]
}
