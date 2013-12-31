// A local cache of an acme.Win's underlying buffer that updates
// itself from the event stream.
//
// For the initial version of win, I want a *windowed* buffer in
// the sense that the buffer stores only a window of the associated
// acme.Win.
//
// Later, I will want to preserve the whole buffer. (Additionally, I will need
// this functionality to implement font highlighting.)

package acmebufs

import (
	"log"
)

// Maintains a slice of an underlying acme buffer
type Winslice struct {
	// Called p
	Offset int				// Offset of Typing into an associated Acme win
	Typing []byte			// UTF8 slice of Acme buffer
}


func New() (*Winslice) {
	return &Winslice{0, make([]byte, 0)}	
}

// assertion: reading [offset, offset + len(Typing)) from Acme land will
// give me the same contents as Typing. Hopefully.

// Do I need anything other than a slice? No?
// Will I need something smarter? Yes. Once I am stashing whole buffers
// to swap in/out

// Adds text t to the Winslice at position p with respect to the
// beginning of the backing Acme buffer. The insertion point needs
// to be within the winslice.
//
// I think that the existing implementation will send newline bounded chunks
// to the child process. This filtering needs to be happening outside of this
// code.
//
// TODO(rjkroege): Clean this up.
//
// Eventually, there will be buffer and ttypair per shell. And a set of win's.
//
func (ws *Winslice) Addtyping(ty []byte, p int) {
	if p < ws.Offset || p > ws.Offset + len(ws.Typing) {
		log.Fatalf("p (%d) !in [ws.Offset: %d, ws.Offset + len)\n", p, ws.Offset)
	}		

	p = p - ws.Offset
	h := ws.Typing[0:p]
	t := ws.Typing[p:]
	tc := make([]byte, len(t))
	copy(tc, t)

	n := append(h[0:p], ty...)
	n = append(n, tc...)
	ws.Typing = n
}

// Advance the offset.
// This function might have to do something more clever
// as I understand the code better.
// I think I have to chop the front off..
func (ws *Winslice) Move(p int) {
	ws.Offset += p	
}





