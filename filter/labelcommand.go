package filter

import (
	"bytes"
)


// Given a buffer of text coming from the pty, remove a label
// command block if present and return the possiblye modified
// buffer and label command.
//
// Like the initial win implementation, punt if the label command
// is not present in its entirety.
func Labelcommand(b []byte) ([]byte, []byte) {
	// find the end of the buffer.
	i := bytes.LastIndex(b, []byte{'\007'})
	if i == -1 {
		return b, nil
	}
	
	al := b[i+1:]
	bl := b[0:i]
	
	if len(bl) < 3 {
		// A full label command won't fit.
		return b, nil
	}

	i = bytes.Index(bl, []byte("\033];"))
	if i == -1 {
		return b, nil
	}
	
	// wrong order..
	il := bl[i+3:]
	bl = bl[0:i]

	var label []byte

	// Idea.
	// A whole selection of interesting commands for winmux could be implemented
	// by simply listening here for out-of-band commentary back from the shell.

	if len(il) > 0 && !bytes.Equal(il, []byte("*9term-hold+")) {
		if !bytes.Contains(il, []byte("/-")) {
			// should add '/-'<command name>
		}
		label = make([]byte, len(il))
		copy(label, il)
	}

	buf := make([]byte, 0, len(al) + len(bl))
	buf = append(buf, bl...)
	buf = append(buf, al...)
	return buf, label
}
