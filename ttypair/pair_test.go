package ttypair

import (
	"github.com/rjkroege/wikitools/testhelpers"
	"testing"
)

func Test_Israw(t *testing.T) {
	tp := New()
	testhelpers.AssertBool(t, false, tp.Israw())

	tp.Setcook(false)
	testhelpers.AssertBool(t, true, tp.Israw())
	tp.Setcook(true)
	testhelpers.AssertBool(t, false, tp.Israw())
}


