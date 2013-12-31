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

