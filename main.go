package main

import (
	"fmt"
	"os"

	"github.com/ALizarazoTellez/Mili/pkg/term"
)

type msg any

type exitMsg struct{}

type model struct {
	char byte
}

func (m *model) update(msg msg) msg {
	if ch, ok := msg.(byte); ok {
		m.char = ch
	}

	if m.char == 'q' {
		return exitMsg{}
	}

	return nil
}

func (m *model) render(buffer [][]byte) {
	for i := range buffer {
		for j := range buffer[i] {
			buffer[i][j] = ' '

			if i == 0 || i == len(buffer)-1 {
				buffer[i][j] = '-'
			}

			if j == 0 || j == len(buffer[i])-1 {
				buffer[i][j] = '|'
			}
		}
	}

	buffer[5][10] = m.char
}

func getchar() byte {
	var char [1]byte

	n, err := os.Stdin.Read(char[:])
	if n != 1 && err != nil {
		panic(err)
	}

	return char[0]
}

func main() {
	term.EnableAltBuffer()
	defer term.DisableAltBuffer()

	state, err := term.GetState()
	if err != nil {
		panic(err)
	}
	defer term.SetState(state)

	err = term.MakeRaw()
	if err != nil {
		panic(err)
	}

	w, h, err := term.Size()
	if err != nil {
		panic(err)
	}

	term.HideCursor()
	defer term.ShowCursor()

	model := model{}

	buffer := make([][]byte, h)
	for i := range buffer {
		buffer[i] = make([]byte, w)
	}

	for {
		model.render(buffer)
		renderBuffer(buffer)

		msg := model.update(getchar())
		if _, ok := msg.(exitMsg); ok {
			break
		}
	}
}

func renderBuffer(buffer [][]byte) {
	for row, rowData := range buffer {
		for col, colData := range rowData {
			term.CursorToHome()
			if row != 0 {
				term.CursorDown(row)
			}

			if col != 0 {
				term.CursorRight(col)
			}
			fmt.Printf("%c", colData)
		}
	}
}
