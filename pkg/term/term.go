package term

import (
	"fmt"
	"os"
	"strconv"

	"golang.org/x/term"
)

var StdinFd = int(os.Stdin.Fd())

const (
	CSI = "\u001b["
)

type State = term.State

func IsTerminal() bool {
	return term.IsTerminal(StdinFd)
}

func GetState() (*State, error) {
	return term.GetState(StdinFd)
}

func SetState(state *State) error {
	return term.Restore(StdinFd, state)
}

func Size() (width, height int, err error) {
	return term.GetSize(StdinFd)
}

func MakeRaw() error {
	_, err := term.MakeRaw(StdinFd)
	return err
}

func CursorToHome() {
	fmt.Print(CSI, "H")
}

func CursorUp(n int) {
	fmt.Print(CSI, strconv.Itoa(n), "A")
}

func CursorDown(n int) {
	fmt.Print(CSI, strconv.Itoa(n), "B")
}

func CursorRight(n int) {
	fmt.Print(CSI, strconv.Itoa(n), "C")
}

func CursorLeft(n int) {
	fmt.Print(CSI, strconv.Itoa(n), "D")
}

func HideCursor() {
	fmt.Print(CSI, "?25l")
}

func ShowCursor() {
	fmt.Print(CSI, "?25h")
}

func EnableAltBuffer() {
	fmt.Print(CSI, "?1049h")
}

func DisableAltBuffer() {
	fmt.Print(CSI, "?1049l")
}
