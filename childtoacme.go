package main
// Ships bytes from the child file source to the acme window.
// TODO(rjk): Move this to its own package.

import (
//	"bytes"
//	"code.google.com/p/goplan9/plan9/acme"
//	"fmt"
	"github.com/rjkroege/winmux/ttypair"
	"github.com/rjkroege/winmux/filter"
	"log"
//	"os"
//	"sync"
	"unicode/utf8"
	//	"code.google.com/p/goplan9/draw"
	//	"image"
//	"flag"
//	"github.com/kr/pty"
//	"os/exec"
	"io"
)

// Must not modify buffer b.
func enhancedlog(b []byte) {

	// display all characters, including control characters.
	nb := make([]byte, len(b))

	// Remember that len(rune) != len(byte). Silly boy.	
	for _, r := range(string(b)) {
		switch r {
		case '\r':
			// stuff that doesn't work...
			nb = append(nb, "\\r"...)
		case '\033':
			nb = append(nb, "\\033"...)
		case '\007':
			nb = append(nb, "\\007"...)
		default:
			nb = append(nb, string(r)...)
		}
	}

	log.Printf("buf: <%s>", string(nb))
}


func childtoacme(q *Q, fd io.Reader, echo *ttypair.Echo) {
	fbuf := make([]byte, 8192 + utf8.UTFMax+1)
	buf := fbuf[0:]
	for {
		nr, er := fd.Read(buf)
		if er == io.EOF {
			log.Printf("EOF on reading from pty: %s", er.Error)
			break
		}
		if er != nil {
			log.Printf("error on reading from pty: %s", er.Error)
			break
		}
		if nr <= 0 {
			continue
		}

		enhancedlog(buf[0:nr])
		// log.Printf("the buffer: <<%s>>", logenhancer(buf[0:nr]))

		b := buf[0:nr]
		b = echo.Cancel(b)
		if len(b) == 0 {
			continue;
		}

		b = filter.Dropcrnl(b)
		if len(b) == 0 {
			continue;
		}

		b = filter.Squashnul(b)
		if len(b) == 0 {
			continue
		}
	
		b, r := filter.Runemodulus(b)
		// TODO(rjk): Remember to do something useful with r
		if len(r) != 0 {
			log.Printf("Runmodulus had a remnant....\n")
		}

		// TODO(rjk): handle label operations here.
		b, label := filter.Labelcommand(b)
		if label != nil {
			q.Win.Name(string(label))
		}

		// TODO(rjk): detect if we have a password prompt, set password true
		// to suppress echo.
		// note need to plumb the call to the ttypair...

		q.Lock()
		err := q.Win.Addr("#%d", q.Tty.Offset)
		if err != nil {
			log.Fatalf("we couldn't handle writing to the Acme. %s\n", err.Error())
		}
		
		_, err = q.Win.Write("data", b)
		if err != nil {
			log.Fatalf("Couldn't write to the acme data: %s\n", err.Error())
		}

		q.Tty.Move(utf8.RuneCount(b))
		q.Unlock()

		// TODO(rjk): Preserve the remmant
		copy(fbuf, r)
		buf = fbuf[len(r):]
	}
}
