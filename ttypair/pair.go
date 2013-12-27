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

// Returns true if t needs to be treated as a raw thinger.
// TODO(rjkroege): Make a type that wraps a io.Reader/Writer?
func (t *Tty) Israw() bool {
//	return (!t.cook || t.password) && !isecho(t.fd0);
	log.Print("Israw\n")
	return false
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

