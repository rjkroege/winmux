/*

    A version of win written in Go. That does terminal multiplexing.

*/

package main

import (
	"fmt"
	"log"
//	"code.google.com/p/goplan9/draw"
//	"image"
	"code.google.com/p/goplan9/plan9/acme"
	"os"
)

func main() {
	fmt.Print("hello from winmux\n");

	log.Print("hello!");

	// take a window id from the command line
	// I suppose it could come from the environment too
	
	log.Print(os.Args[0])

	var win *acme.Win
	var err error

	// TODO(rjkroege): look up a window by name if an argument is provided
	if len(os.Args) > 1 {
		log.Fatal("write some code to lookup window by name and connect")
	} else {
		win,err = acme.New()
	}
	if err != nil {
		log.Fatal("can't open the window? ", err.Error())
	}

	win.Fprintf("body", "hi rob")
	win.CloseFiles()
	fmt.Print("bye\n")
}

