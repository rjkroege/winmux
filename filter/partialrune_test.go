package filter

import (
	"github.com/rjkroege/wikitools/testhelpers"
	"testing"
)

func Test_Runemodulus(t *testing.T) {

	b, r := Runemodulus([]byte("abcd"))
	testhelpers.AssertString(t, "abcd", string(b))
	testhelpers.AssertInt(t, 0, len(r))

	// 日本語  <- 3 Asian letters.
	s := "\xe6\x97\xa5\xe6\x9c\xac\xe8\xaa\x9e"
	sb := []byte(s)

	b, r = Runemodulus(sb)
	testhelpers.AssertString(t, s, string(b))
	testhelpers.AssertInt(t, 0, len(r))

	b, r = Runemodulus(sb[:len(sb)-1])
	testhelpers.AssertString(t, s[0:6], string(b))
	testhelpers.AssertInt(t, 2, len(r))
	testhelpers.AssertInt(t, int(int(r[0])), '\xe8')
	testhelpers.AssertInt(t, int(int(r[1])), '\xaa')

	b, r = Runemodulus(sb[:len(sb)-2])
	testhelpers.AssertInt(t, 1, len(r))
	testhelpers.AssertString(t, s[0:6], string(b))
	testhelpers.AssertInt(t, int(r[0]), '\xe8')

}
