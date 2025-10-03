//go:build !gio
// +build !gio

package main

import "fmt"

// runGUI is a stub when the project is built without the 'gio' build tag.
func runGUI(apiKey string) {
	fmt.Println("GUI not available in this build.")
	fmt.Println("To enable the Gio GUI, install the gioui.org modules and build/run with:  ")
	fmt.Println("  go run -tags gio . -gui")
}
