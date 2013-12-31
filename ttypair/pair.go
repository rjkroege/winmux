/*
	Manages a tty pair.
*/


package ttypair

import (
	// TODO(rjkroege): suck in the bsd pty
	"bytes"
	"log"
	"code.google.com/p/goplan9/plan9/acme"
)

type Tty struct  {
	bytes.Buffer
	cook bool
	password bool
	fd0 int	
}

// Creates a Tty object
func New() (*Tty) {
	return &Tty{cook: true, password: false, fd0: -1}
}

// Returns true if t needs to be treated as a raw tty.
func (t *Tty) Israw() bool {
	log.Print("Israw\n")
	// TODO(rjkroege): Pull in isecho.
	return (!t.cook || t.password) /* && !isecho(t.fd0) */;
}

// Deletes characters from the buffer etc
func (t *Tty) Delete(e *acme.Event) int {
	log.Print("Delete\n")
	return 1
}

// Ships n backspaces to the child.
func (t *Tty) Sendbs(n int) {
	log.Printf("Sendbs %d\n", n)
}

func (t *Tty) Setcook(b bool) {
	t.cook = b;
	log.Printf("Setcook to %b\n", b)
}

// Writes the provided buffer to the associated file descriptor.
// Usually used to ship a 0x7F to the remote (so far)
// probably care about errors...
func (t *Tty) UnbufferedWrite(b []byte) error {
	log.Print("attempting to write a delete to the remote\n")
	return nil
}

/*
	In win, the buffer is a window onto the larger Acme window.
	Much of the complexity is in supporting that `typing` buffer is
	a small window into a larger buffer.

	I will need a complete buffer to support muxing / autosave.

	Perhaps I should accept this now? The right way is to maintain a
	parallel buffer and apply edits to it. 

	The low road way is to just re-read the buffer.This can be replaced
	with something clever where I collect the edits?

	Buffer management is going to go badly with the passwordy stuff.
	What to do next...

	Let's get it working first in the existing way. I need a buffer class where we
	accumulate typing
*/

func (t *Tty) addtype(e *acme.Event) {
	// I need to manage a buffer.
	log.Print("addtype... do the buffer management\n")
}

// Add typing to the buffer or do a bypass write as necessary
// TODO(rjkroege): This is not in the right place.
func (t *Tty) Type(e *acme.Event) {
	log.Printf("should add the typing to the buffer?\n")

	// what about case where amount added is too large to be in event?

	if e.Nr > 0 {
		// Call addtype..
		t.addtype(e)
	} else {
		log.Fatal("you've not handled the case where you need to read from acme\n")
	}

	if t.Israw() {
		// This deletes the character typed if we have set israw so that
		// raw mode works properly.
		log.Printf("unsupported raw mode\n");
//		n = sprint(buf, "#%d,#%d", e->q0, e->q1);
//		fswrite(afd, buf, n);
//		fswrite(dfd, "", 0);
//		q.p -= e->q1 - e->q0;
	}
	t.sendtype()
	if len(e.Text) > 0 && e.Text[len(e.Text) - 1] == '\n' {
		// Not really clear to me what this is for.
		t.cook = true;
	}
}


func (t *Tty) sendtype() {
	log.Print("write sendtype\n")
}

// Return some kind of count of something in the in-progress typing buffer.
func (t *Tty) Ntyper() int {
	// TODO(rjkroege): Write me.
	return 1
}