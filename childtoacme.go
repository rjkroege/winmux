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
	buf := make([]byte, 8192 + utf8.UTFMax+1)
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

		buf = echo.Cancel(buf)
		if len(buf) == 0 {
			continue;
		}

		// TODO(rjkroege): Convert this function
		buf = filter.Dropcrnl(buf);
		if len(buf) == 0 {
			continue;
		}

		// Must drive this downwards.
		dumpoutput(b)

	}
}


func childtoacme_old(q *Q, fd io.Reader) {
	
	buf := make([]byte, 8192 + utf8.UTFMax+1)

	for {
		/* Let typing have a go -- maybe there's a rubout waiting. */
		// yield();
		
		// Need to update the length or buffer will not be right length.
		c, error := fd.Read(buf)
		if error != nil {
			// Probabaly the command failed?
			// TODO(rjkroege): Figure out what to do here.
			log.Fatal("Read failed: %s\n", error.Error)
		}
		if c == 0 {
			continue
		}
		log.Printf("go routine doing thing...")

		// Temporary. TODO(rjk): Re-process
		dumpoutput(buf[0:c])

		// TODO(rjkroege): Convert this function
		//n = echocancel(buf+npart, n);
		//if(n == 0)
		//	continue;

		// TODO(rjkroege): Convert this function
		// n = dropcrnl(buf+npart, n);
		// if(n == 0)
		//	continue;

		/* squash NULs */
		//
		//s = memchr(buf+npart, 0, n);
		//if(s){
		//	for(t=s; s<buf+npart+n; s++)
		//		if(*t = *s)	/* assign = */
		//			t++;
		//	n = t-(buf+npart);
		//}

		//n += npart;

		// This section looks through what we're given and makes sure
		// that we end on an even rune boundary.
		/* hold on to final partial rune */
		//npart = 0;
		//while(n>0 && (buf[n-1]&0xC0)){
		//	--n;
		//	npart++;
		//	if((buf[n]&0xC0)!=0x80){
		//		if(fullrune(buf+n, npart)){
		//			w = chartorune(&r, buf+n);
		//			n += w;
		//			npart -= w;
		//		}
		//		break;
		//	}
		//}

		// password filtering
		//if(n > 0){
//			memmove(hold, buf+n, npart);
//			buf[n] = 0;
//			n = label(buf, n);
//			buf[n] = 0;
//			
//			// clumsy but effective: notice password
//			// prompts so we can disable echo.
//			password = 0;
//			if(cistrstr(buf, "password") || cistrstr(buf, "passphrase")) {
//				int i;
//				
//				i = n;
//				while(i > 0 && buf[i-1] == ' ')
//					i--;
//				password = i > 0 && buf[i-1] == ':';
//			}
//
//			q.Lock();
//			m = sprint(x, "#%d", q.p);
//			if(fswrite(afd, x, m) != m){
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
//			if(fswrite(dfd, buf, n) != n)
//				error("stdout writing body");
//			/* Make sure acme scrolls to the end of the above write. */
//			if(fswrite(dfd, nil, 0) != 0)
//				error("stdout flushing body");
//			q.p += nrunes(buf, n);
//			q.Unlock();
//			memmove(buf, hold, npart);
//		}
	}

	log.Printf("leaving childtoacme")
}
