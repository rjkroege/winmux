/*
	Manages a tty pair.
*/


package ttypair

import (
	// TODO(rjkroege): suck in the bsd pty
	"bytes"
	"log"
	"code.google.com/p/goplan9/plan9/acme"
	"github.com/rjkroege/winmux/acmebufs"
)

type Tty struct  {
	acmebufs.Winslice
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
// Either a single delete character to stop the remote or a single
// command line for the remote shell to execute.
// TODO(rjkroege): Send the provided buffer off to the child process.
func (t *Tty) UnbufferedWrite(b []byte) error {
	log.Println("UnbufferedWrite: ", string(b))
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

// Adds typing to the buffer associated with this pair at position p0.
func (t *Tty) addtype(typing []byte, p0 int, fromkeyboard bool) {
	log.Print("addtype... do the buffer management\n")
	if bytes.Index(typing, []byte{3, 0x7}) != -1 {
		t.Reset()
	}
	t.Addtyping(typing, p0)
}

// Add typing to the buffer or do a bypass write as necessary
// TODO(rjkroege): This is not in the right place.
func (t *Tty) Type(e *acme.Event) {
	log.Printf("should add the typing to the buffer?\n")

	// what about case where amount added is too large to be in event?

	if e.Nr > 0 {
		// TODO(rjkroege): Conceivably, I am not shifting the offset enough.
		t.addtype(e.Text, e.Q0, e.C1 == 'K' /* Verify this test. */)
	} else {
		log.Fatal("you've not handled the case where you need to read from acme\n")
		// TODO(rjkroege): Write the acme fetcher...
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

// This is sendtype !raw. 
// TODO(rjkroege): Write sendtype_raw too.
func (t *Tty) sendtype() {
	log.Print("write sendtype\n")

	// raw and cooked mode are interleaved. Write cooked mode
	// aside: we should be removing the typed characters in acme right 
	// because otherwise the echo will insert them twice... (this block of code)

	typebreaks := bytes.Split(t.Typing, []byte{ '\n', 0x04 })
	for _,  s := range typebreaks[0:len(typebreaks)-1] {
		// Skip the last one because it's the text *following* the newline.
		echoed(s)
		t.UnbufferedWrite(s)	// Send to the child program
	}
	
	// Does this mean that the store backing it grows indefinitely?
	// I think yes. I should copy.
	t.Typing = typebreaks[len(typebreaks)-1]
}

// Inserts the provided buffer into Acme.
func echoed(s []byte) {
	// TODO(rjkroege): Write me
	log.Print("echoing back...\n")

	/*
		Only one thread can be writing to the acme buffer at
		a time. The existing win implementation uses a lock.
		Would I prefer to not be lock based and simply have
		a single thread that processes a series of messages
		to update the acme.

		we therefore might have the following threads:

		*  listener for events
		*  waiting on channel, updates acme
		*  listener for output from rc

		which obviates the need for locks

		I will use a lock.
	*/
	// TODO(rjkroege): I need the win construct.
	
}
