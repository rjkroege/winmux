// A local cache of an acme.Win's underlying buffer that updates
// itself from the event stream.
//
// For the initial version of win, I want a *windowed* buffer in
// the sense that the buffer stores only a window of the associated
// acme.Win.
//
// Later, I will want to preserve the whole buffer. (Additionally, I will need
// this functionality to implement font highlighting.)

// Bikeshedding: this class is not named well. It is the last line of the
// acme win that has not been sent to the shell. Hence I see the point
// of calling it "Typing" in the original code.

package acmebufs

import (
	"fmt"
	"log"
)

// Maintains a slice of an underlying acme buffer
type Winslice struct {
	// Called p
	Offset int    // Offset of Typing into an associated Acme win
	Typing []byte // UTF8 slice of Acme buffer
}

func (ws *Winslice) String() string {
	return string(ws.Typing)
}

func New() *Winslice {
	return &Winslice{0, make([]byte, 0)}
}

// Resets the winslice to no longer contain text.
// TODO(rjkroege): Make sure that this is right. In particular,
// that we don't toss the offset.
func (ws *Winslice) Reset() {
	ws.Typing = ws.Typing[0:0]
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
	log.Println("Winslice.Addtyping")
	if p < ws.Offset || p > ws.Offset+len(ws.Typing) {
		panic(fmt.Sprintf("p (%d) !in [ws.Offset: %d, ws.Offset + len: %d)", p, ws.Offset, ws.Offset + len(ws.Typing)))
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

// Move the offset.
// This function might have to do something more clever
// as I understand the code better. It might have to chop the
// front off the typing.
func (ws *Winslice) Move(p int) {
	ws.Offset += p
}

// Removes the typing specified by the range returning the number
// of characters deleted outside of the range (the amount that we
// should adjust the offset following the Delete.)
// TODO(rjkroege): Should this method update the offset?
func (ws *Winslice) Delete(q0, q1 int) int {
	log.Printf("Delete %d, %d\n", q0, q1)
	var lq0, lq1, d int
	switch {
	case q0 < ws.Offset && q1<= ws.Offset:
		log.Printf("Delete branch a")
		return q1 - q0
	case ws.Inslice(q1) && q0 < ws.Offset && q1 > ws.Offset:
		log.Printf("Delete branch b")
		lq0 = 0
		lq1 = q1 - ws.Offset
		d = lq1
	case ws.Inslice(q0) && ws.Inslice(q1) && q0 >= ws.Offset && q1 > ws.Offset:
		log.Printf("Delete branch c")
		lq0 = q0 - ws.Offset
		lq1 = q1 - ws.Offset
		d = lq1 - lq0
	default:
		d = 0
		panic("Delete went wrong")
	}
	
	copy(ws.Typing[lq0:], ws.Typing[lq1:])
	ws.Typing = ws.Typing[0:len(ws.Typing) - d]
	return (q1 - q0) - d
}

// Is the provided position q0 in the Acme buffer before the start of
// the slice. In particular: the ws is usually the last incomplete line
// of text yet to be delivered to the shell and Offset is its start.
func (ws *Winslice) Beforeslice(q0 int) bool {
	return q0 < ws.Offset
}

// Returns true if the given address is within the slice.
func (ws *Winslice) Inslice(q0 int) bool {
	return q0 <= ws.Offset+len(ws.Typing)
}

// Returns true if the given position is at the end or beyond the
// the slice.
// I think. It's not quite clear what this is testing...
func (ws *Winslice) Afterslice(q0, n int) bool {
	return q0 >= ws.Offset+n
}

func (ws *Winslice) Ntyper() int {
	return len(ws.Typing)
}

// Use this for logging.
// NB: ws.Offset corresponds to "p" and len(ws.Typing) to ntyper
// TODO(rjkroege): Offset could be called p
// TODO(rjkroege): someday. refactor this nicely.
func (ws *Winslice) Extent() (int, int) {
	return ws.Offset, ws.Offset + len(ws.Typing)
}
