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
	"os"
	"io"
)

// Temporary code, to emit to stdout.
func dumpoutput(buf []byte) {
	_, error := os.Stdout.Write(buf)
	log.Printf("dumpoutput?")
	if error != nil {
		log.Fatal("Couldn't copy to Stdout: %s", error.Error())
	}
	// return buf[0:cap(buf)]
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

		b := buf[0:nr]
		b = echo.Cancel(b)
		if len(b) == 0 {
			continue;
		}

		b = filter.Dropcrnl(b)
		if len(b) == 0 {
			continue;
		}

		// TODO(rjkroege): HERE. write this one.
		b = filter.Squashnul(b)
		if len(b) == 0 {
			continue
		}
	
		b, r := filter.Runemodulus(b)
		// TODO(rjk): Remember to do something useful with r
		log.Printf("Runmodulus had a remnant....\n")		
	
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


func childtoacme_old(q *Q, fd io.Reader) {
//
//			q.Lock();

// set insertion point
//			m = sprint(x, "#%d", q.p);
//			if(fswrite(afd, x, m) != m){
// clean up if something went wrong.
//				fprint(2, "stdout writing address %s: %r; resetting\n", x);
//				if(fswrite(afd, "$", 1) < 0)
//					fprint(2, "reset: %r\n");
//				fsseek(afd, 0, 0);
//				m = fsread(afd, x, sizeof x-1);
//				if(m >= 0){
//					x[m] = 0;
//					q.p = atoi(x);
//				}
//			}
// insert the actual text.
//			if(fswrite(dfd, buf, n) != n)
//				error("stdout writing body");
//			/* Make sure acme scrolls to the end of the above write. */
// scroll to the bottom
//			if(fswrite(dfd, nil, 0) != 0)
//				error("stdout flushing body");
//			q.p += nrunes(buf, n);
//			q.Unlock();
//			memmove(buf, hold, npart);
//		}

	log.Printf("leaving childtoacme")
}
