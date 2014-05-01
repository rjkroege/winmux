package filter

import (
	"unicode/utf8"
)

// Given a byte slice, divides it into the prefix that consists only of
// valid unicode characters and the remaining bytes that are not
// part of a valid unicde sequence.
func Runemodulus(buf []byte) (valid []byte, rem []byte) {
	var i int
	for i = len(buf) - 1; i > 0 && (buf[i] & 0xC0) == 0x80; i-- { }
	if buf[i] & 0xC0 == 0xC0 && !utf8.Valid(buf[i:]) {
		rem = buf[i:]
		valid = buf[0:i]
		return
	}

	valid = buf
	rem = make([]byte, 0)
	return
}
