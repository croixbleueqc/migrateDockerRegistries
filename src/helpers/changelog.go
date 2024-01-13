// nxrmuploader
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/helpers/changelog.go
// Original timestamp: 2023/12/31 14:45

package helpers

import "fmt"

func ChangeLog() {
	//fmt.Printf("\x1b[2J")
	fmt.Printf("\x1bc")

	CenterPrint("CHANGELOG")
	fmt.Println()
	CenterPrint("=========")
	fmt.Println()
	fmt.Println()

	fmt.Print(`
VERSION			DATE			COMMENT
-------			----			-------
1.00.00			2024.01.11		initial version
`)
}
