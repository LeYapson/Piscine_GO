package main

import "github.com/01-edu/z01"

type Door struct {
	state bool
}

const (
	OPEN  = false
	CLOSE = true
)

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func OpenDoor(ptrDoor *Door) bool {
	PrintStr("Door Opening...")
	ptrDoor.state = OPEN
	return true
}

func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Closing...")
	ptrDoor.state = CLOSE
	return true
}

func IsDoorOpen(ptrDoor *Door) bool {
	PrintStr("is the Door opened ?")
	return ptrDoor.state
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	return ptrDoor.state
}

func main() {
	door := &Door{}

	OpenDoor(door)
	z01.PrintRune('\n')
	if IsDoorClose(door) {
		OpenDoor(door)
		z01.PrintRune('\n')
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
		z01.PrintRune('\n')
	}
	if !door.state {
		CloseDoor(door)
		z01.PrintRune('\n')
	}
}
