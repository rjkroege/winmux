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

// Adds typing to the buffer associated with this pair at position p0.
func (t *Tty) addtype(typing []byte, p0 int, fromkeyboard bool) {
	log.Println("Tty.addtype")
	if bytes.Index(typing, []byte{3, 0x7}) != -1 {
		log.Println("Tty.addtype: resetting")
		t.Reset()
	}
	t.Addtyping(typing, p0)
}

// Add typing to the buffer or do a bypass write as necessary
// TODO(rjkroege): This is not in the right place.
func (t *Tty) Type(e *acme.Event) {
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
		log.Fatal("unsupported raw mode\n");
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
// TODO(rjkroege): Write sendtype_raw or modify this function to do raw mode.
// TODO(rjkroege): this is buffer-oriented. maybe move into winslice?
func (t *Tty) sendtype() {
	// raw and cooked mode are interleaved. Write cooked mode
	// aside: we should be removing the typed characters in acme right 
	// because otherwise the echo will insert them twice... (this block of code)

	ty := t.Typing;
	mutated := false	
	for p := bytes.IndexAny(ty, "\n\004"); p >= 0; p = bytes.IndexAny(ty, "\n\004") {
		s := ty[0:p+1]
		echoed(s)
		t.UnbufferedWrite(s)	// Send to the child program
		t.Move(len(s))
		mutated = true
 		ty = ty[p+1:]
	}

	// Copy the remaining text to a new slice so that the old backing can
	// get garbage collected.
	if (mutated) {
		t.Typing = make([]byte, len(ty))
		copy(t.Typing, ty)
	}
}

// Inserts the provided buffer into Acme from sendtype.
// TODO(rjkroege): Write echod
func echoed(s []byte) {
	log.Print("echoed")		
}
