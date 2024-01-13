// migrateDockerRegistries
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/helpers/terminal.go

package helpers

import (
	"fmt"
	"github.com/jwalton/gchalk"
	"syscall"
	"unsafe"
)

var PlainOutput = false

// COLOR FUNCTIONS
// ===============
func Red(sentence string) string {
	return fmt.Sprintf("%s", gchalk.WithBrightRed().Bold(sentence))
}

func Green(sentence string) string {
	return fmt.Sprintf("%s", gchalk.WithBrightGreen().Bold(sentence))
}

func White(sentence string) string {
	return fmt.Sprintf("%s", gchalk.WithBrightWhite().Bold(sentence))
}

func Yellow(sentence string) string {
	return fmt.Sprintf("%s", gchalk.WithBrightYellow().Bold(sentence))
}

func Blue(sentence string) string {
	return fmt.Sprintf("%s", gchalk.WithBrightBlue().Bold(sentence))
}

// FIXME : Normal() is the same as White()
func Normal(sentence string) string {
	return fmt.Sprintf("%s", gchalk.WithWhite().Bold(sentence))
}

// TERMINAL FUNCTIONS
func GetTerminalSize() (int, int) {
	var size struct {
		rows    uint16
		cols    uint16
		xpixels uint16
		ypixels uint16
	}
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(syscall.Stdin), syscall.TIOCGWINSZ, uintptr(unsafe.Pointer(&size)))
	if err != 0 {
		return 0, 0
	}
	return int(size.cols), int(size.rows)
}

func CenterPrint(text string) {
	termWidth, _ := GetTerminalSize()
	padding := (termWidth - len(text)) / 2
	fmt.Printf("%s[%dC%s", terminalEscape, padding, text)
}
