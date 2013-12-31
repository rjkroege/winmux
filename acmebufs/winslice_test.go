package acmebufs

import (
	"github.com/rjkroege/wikitools/testhelpers"
	"testing"
)

func Test_Move(t *testing.T) {
	ws := New()

	testhelpers.AssertInt(t, 0, ws.Offset)
	ws.Move(1)
	testhelpers.AssertInt(t, 1, ws.Offset)
	ws.Move(-1)
	testhelpers.AssertInt(t, 0, ws.Offset)
}

func Test_Addtyping(t *testing.T) {
	ws := New()
	ws.Move(2)

	ws.Addtyping([]byte{'a'}, 2)
	testhelpers.AssertString(t, "a", string(ws.Typing))

	ws.Addtyping([]byte{'b'}, 3)
	testhelpers.AssertString(t, "ab", string(ws.Typing))

	ws.Addtyping([]byte{'c'}, 4)
	testhelpers.AssertString(t, "abc", string(ws.Typing))

	testhelpers.AssertInt(t, 2, ws.Offset)
	ws.Addtyping([]byte{'X'}, 2)
	testhelpers.AssertString(t, "Xabc", string(ws.Typing))

	ws.Addtyping([]byte{'Y'}, 3)
	testhelpers.AssertString(t, "XYabc", string(ws.Typing))
	
}

