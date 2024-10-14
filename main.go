package main

import (
	"fmt"
	"os"

	"github.com/ALizarazoTellez/Mili/pkg/term"
)

func do() {
	term.CursorToHome()
	term.HideCursor()
	defer term.ShowCursor()

	fmt.Println("Hello World!")
	w, h, err := term.Size()
	if err != nil {
		panic(err)
	}
	fmt.Println("Size is:", w, h)

	ch := getchar()
	fmt.Print(ch)
	fmt.Printf("%c", ch)
	fmt.Println(getchar())
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

	do()
}
