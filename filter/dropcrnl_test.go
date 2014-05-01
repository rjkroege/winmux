package filter

import (
	"github.com/rjkroege/wikitools/testhelpers"
	"testing"
)

func Test_Dropcrnl(t *testing.T) {
	testhelpers.AssertString(t, "", string(Dropcrnl([]byte(""))))
	testhelpers.AssertString(t, "h", string(Dropcrnl([]byte("h"))))
	testhelpers.AssertString(t, "hello", string(Dropcrnl([]byte("hello"))))
	testhelpers.AssertString(t, "hello\nthere", string(Dropcrnl([]byte("hello\r\nthere"))))
	testhelpers.AssertString(t, "hello\nthere", string(Dropcrnl([]byte("hello\nthere"))))
}
