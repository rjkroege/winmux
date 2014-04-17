package filter

import (
	"github.com/rjkroege/wikitools/testhelpers"
	"testing"
)

func Test_Squashnul(t *testing.T) {
	testhelpers.AssertString(t, "", string(Squashnul([]byte(""))))
	testhelpers.AssertString(t, "h", string(Squashnul([]byte("h"))))
	testhelpers.AssertString(t, "hello", string(Squashnul([]byte("hello"))))
	testhelpers.AssertString(t, "hello", string(Squashnul([]byte("he\x00\x00llo\x00\x00"))))
}
