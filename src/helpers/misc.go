// migrateDockerRegistries
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/helpers/misc.go

package helpers

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

const (
	terminalEscape = "\x1b"
)

// CustomError implements the error interface
type CustomError struct {
	Message string
}

func (e CustomError) Error() string {
	return e.Message
}

// reference: https://gist.github.com/jlinoff/e8e26b4ffa38d379c7f1891fd174a6d0, the getPassword2.go
func GetPassword(prompt string) string {
	// Get the initial state of the terminal.
	initialTermState, e1 := terminal.GetState(syscall.Stdin)
	if e1 != nil {
		panic(e1)
	}

	// Restore it in the event of an interrupt.
	// CITATION: Konstantin Shaposhnikov - https://groups.google.com/forum/#!topic/golang-nuts/kTVAbtee9UA
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	go func() {
		<-c
		_ = terminal.Restore(syscall.Stdin, initialTermState)
		os.Exit(1)
	}()

	// Now get the password.
	fmt.Print(prompt)
	p, err := terminal.ReadPassword(syscall.Stdin)
	fmt.Println("")
	if err != nil {
		panic(err)
	}

	// Stop looking for ^C on the channel.
	signal.Stop(c)

	// Return the password as a string.
	return string(p)
}

// NUMBER FORMATTING FUNCTIONS
// ===========================

// This function was originally written in 1993, in C, by my friend Jean-François Gauthier (jief@brebis.dyndns.org)
// I've ported it in C# in 2011. It is then a third iteration of this function
// This function transforms a multi-digit number in International Notation; thus 1234567 becomes 1,234,567
func SI(nombre uint64) string {
	var strN string
	var strbR bytes.Buffer
	var nLen, nPos int

	strN = strconv.FormatUint(nombre, 10)
	strN = ReverseString(strN)
	nLen = len(strN)

	for nPos < nLen {
		if nPos != 0 && nPos%3 == 0 {
			strbR.WriteString(",")
			strbR.WriteString(string(strN[nPos]))
		} else {
			strbR.WriteString(string(strN[nPos]))
		}
		nPos++
	}

	strN = strbR.String()
	strN = ReverseString(strN)

	return strN
}

// OTHER FUNCTIONS
// ===============

// This function takes a string and returns its reverse
// Thus, "12345" becomes "54321"
func ReverseString(inputStr string) string {
	runes := []rune(inputStr)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
