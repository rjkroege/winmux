package ttypair

import (
	"github.com/rjkroege/wikitools/testhelpers"
	"testing"
)


func Test_EchoedType(t *testing.T) {
	echo := Makecho()
	s := []byte("hello")
	
	testhelpers.AssertInt(t, 0, len(echo.childbound))

	echo.echoed(s)
	s[0] = 'H'

	testhelpers.AssertInt(t, 1, len(echo.childbound))
	r := <-echo.childbound
	testhelpers.AssertString(t, "hello", string(r))
	testhelpers.AssertInt(t, 0, len(echo.childbound))
}


func Test_Cancel_Wholeword(t *testing.T) {
	echo := Makecho()
	echo.echoed([]byte("hello"))

	s := []byte("h")
	r := echo.Cancel(s)
	testhelpers.AssertInt(t, 0, len(r))
	testhelpers.AssertString(t, "ello", string(echo.oldest))
	
	s = []byte("e")
	r = echo.Cancel(s)
	testhelpers.AssertInt(t, 0, len(r))
	testhelpers.AssertString(t, "llo", string(echo.oldest))
}

func Test_Cancel_Subword(t *testing.T) {
	echo := Makecho()
	echo.echoed([]byte("hello\n"))

	s := []byte("hello\r\nworld")
	r := echo.Cancel(s)
	testhelpers.AssertString(t, "world", string(r))
	if echo.oldest != nil {
		t.Errorf("Test_Cancel_Subword: failed to reset oldest: <%s>", string(echo.oldest))
	}
}

func Test_Cancel_Backspace(t *testing.T) {
	echo := Makecho()
	echo.echoed([]byte("helloworld"))

	s := []byte("hello\x08 \x08worldyay")
	r := echo.Cancel(s)
	testhelpers.AssertString(t, "yay", string(r))
	if echo.oldest != nil {
		t.Error("Test_Cancel_Subword: failed to reset oldest")
	}
}

func Test_Cancel_CharacterChange(t *testing.T) {
	echo := Makecho()
	echo.echoed([]byte("hello"))

	s := []byte("helloworld")
	r := echo.Cancel(s)
	testhelpers.AssertString(t, "world", string(r))
	if echo.oldest != nil {
		t.Error("Test_Cancel_Subword: failed to reset oldest: ", string(echo.oldest))
	}
}

func Test_Cancel_Terminating_NL_Not_Cancelled(t *testing.T) {
	echo := Makecho()
	echo.echoed([]byte("ls\n"))

	s := []byte("ls\r\n")
	r := echo.Cancel(s)
	testhelpers.AssertString(t, "", string(r))
}
