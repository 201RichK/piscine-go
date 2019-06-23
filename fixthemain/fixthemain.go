package main

import (
	"github.com/01-edu/z01"
)

const OPEN = false
const CLOSE = true

type Door struct {
	state bool
}

func PrintStr(str string) {
	arrayRune := []rune(str)
	for _, s := range arrayRune {
		z01.PrintRune(s)
	}
}

func OpenDoor(Door *Door) {
	PrintStr("Door Opening...\n")
	Door.state = OPEN
}

func CloseDoor(Door *Door) bool {
	PrintStr("Door closing...\n")
	Door.state = CLOSE
	return true
}

func IsDoorOpen(Door Door) bool {
	PrintStr("Door is open ?\n")
	return Door.state
}

func IsDoorClose(Door Door) bool {
	PrintStr("Door is close ?\n")
	return Door.state
}

func main() {
	door := Door{}

	OpenDoor(&door)
	if IsDoorClose(door) {
		OpenDoor(&door)
	}
	if IsDoorOpen(door) {
		CloseDoor(&door)
	}
	if door.state == OPEN {
		CloseDoor(&door)
	}
}
