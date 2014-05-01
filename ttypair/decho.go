// It is conceivable that this is mis-named.
// Oh well.
package ttypair

import (
	"log"
)

type Echo struct {
	childbound chan []byte
	oldest []byte
}

func Makecho() *Echo {
	return &Echo{make(chan []byte, 100), nil}
}

// Stashes the slice s so that later we can we can compare it with the response
// from the child process and discard the echo provided by the child process.
// Makes a copy so that we do not acquire a dependency on the slice backing.
func (echos *Echo) echoed(s []byte) {
	log.Print("echoed")

	e := make([]byte, len(s))
	copy(e, s)
	echos.childbound <- e
}

// compare the given string that has come from the child
// process with the previously recorded string sent to the
// child process and delete characters that are part of
// what we have echoed

// Remove commands that have been previously been recorded
// as sent to the child process because they are already in the acme
// buffer.
func (e *Echo) Cancel(p []byte) []byte {
	if e.oldest == nil && len(e.childbound) > 0 {
		e.oldest = <- e.childbound
	}

	if e.oldest != nil {
		var i, j int
		for i, j = 0, 0; i < len(p) && j < len(e.oldest); i++ {
			if e.oldest[j] == p[i] {
				j++
				continue
			} else if e.oldest[j] == '\n' && p[i] == '\r' {
				continue
			} else if (p[i] == 0x08) {
				if i+2 <= len(p) && p[i+1] == ' ' && p[i+2] == 0x08 {
					i += 2
				}
				continue
			}
			break
		}
		p = p[i:]
		if len(e.oldest[j:]) == 0 {
			e.oldest = nil						
		} else {
			e.oldest = e.oldest[j:]
		}
	}
	return p
}
