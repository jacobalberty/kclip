package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

type controlCollection map[string]Control

var cList = controlCollection{
	"print": {"\033[5i%s\033[4i", nil, nil},
	"tmux":  {"\033Ptmux;%s\033\\", func(s string) string { return strings.ReplaceAll(s, "\033", "\033\033") }, nil},
	"osc52": {"\033]52;c;%s\a", nil, base64.StdEncoding.EncodeToString},
}

type Control struct {
	CodeF string
	EF    func(string) string
	WF    func(src []byte) string
}

type controlEncoder struct {
	CList []Control
	W     io.Writer
}

func (c controlEncoder) Write(p []byte) (n int, err error) {
	var buf bytes.Buffer
	n, err = buf.Write(p)
	if err != nil {
		return
	}

	cTmp := "%s"
	for _, cc := range c.CList {
		if cc.EF != nil {
			cTmp = cc.EF(cTmp)
		}
		if cc.WF != nil {
			buf = *bytes.NewBufferString(cc.WF(buf.Bytes()))
		}
		cTmp = fmt.Sprintf(cc.CodeF, cTmp)
	}

	_, err = c.W.Write([]byte(fmt.Sprintf(cTmp, buf.String())))
	return
}

func (c controlCollection) Current(w io.Writer) io.Writer {
	ce := controlEncoder{W: w}
	ce.CList = append(ce.CList, c["print"])
	//ce.CList = append(ce.CList, c["osc52"]) // TODO: Decide how to pick osc52 instead of print control
	if len(os.Getenv("TMUX")) > 0 {
		ce.CList = append(ce.CList, c["tmux"])
	}
	return ce
}

func kclipCopy(dst io.Writer, src io.Reader) (int64, error) {
	pads := cList.Current(dst)
	return io.Copy(pads, src)
}

func getDA() (Pp, Pv, Pc string) {
	var attr string
	var reading bool

	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	defer exec.Command("stty", "-F", "/dev/tty", "echo").Run()
	fmt.Print("\033[>c")

	var b []byte = make([]byte, 1)
	if err := syscall.SetNonblock(0, true); err != nil {
		panic(err)
	}
	defer syscall.SetNonblock(0, false)
	f := os.NewFile(0, "stdin")
	if err := f.SetDeadline(time.Now().Add(500 * time.Millisecond)); err != nil {
		panic(err)
	}
rLoop:
	for {
		_, err := f.Read(b)
		if err != nil {
			break
		}
		switch b[0] {
		case '[':
		case '>':
		case '\033':
			reading = true
		case 'c':
			break rLoop
		default:
			if reading {
				attr += string(b)
			}
		}
	}
	tmp := strings.Split(attr, ";")
	if len(tmp) < 3 {
		return "", "", ""
	}
	return tmp[0], tmp[1], tmp[2]
}
