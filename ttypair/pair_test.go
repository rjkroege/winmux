package ttypair

import (
	"9fans.net/go/acme"
	"github.com/rjkroege/wikitools/testhelpers"
	"testing"
)

func Test_Israw(t *testing.T) {
	tp := New(new(mockttyfd), Makecho())
	testhelpers.AssertBool(t, false, tp.Israw())

	tp.Setcook(false)
	testhelpers.AssertBool(t, true, tp.Israw())
	tp.Setcook(true)
	testhelpers.AssertBool(t, false, tp.Israw())
}

// TODO(rjkroege): Make error testing more robust.
type mockttyfd struct {
	writes [][]byte
}

func (mt *mockttyfd) Write(b []byte) (int, error) {
	mt.writes = append(mt.writes, b)
	return len(b), nil
}

func Test_addtype(t *testing.T) {
	tp := New(new(mockttyfd), Makecho())

	tp.addtype([]byte("hello"), 0, false)
	testhelpers.AssertString(t, "hello", tp.String())

	// addtype is doing the wrong thing...
	tp.addtype([]byte{3}, len("hello"), false)
	testhelpers.AssertString(t, "hello", tp.String())

	// addtype is doing the wrong thing...
	tp.addtype([]byte{3}, len("hello_"), true)
	testhelpers.AssertString(t, "", tp.String())

}

func Test_Sendtype(t *testing.T) {
	mock := &mockttyfd{make([][]byte, 0, 10)}
	tp := New(mock, Makecho())

	tp.addtype([]byte("hello\nbye"), 0, false)
	tp.Sendtype()

	testhelpers.AssertInt(t, 1, len(mock.writes))
	testhelpers.AssertString(t, "hello\n", string(mock.writes[0]))
	testhelpers.AssertString(t, "bye", string(tp.Typing))
}

func Test_SendtypeOnechar(t *testing.T) {
	mock := &mockttyfd{make([][]byte, 0, 10)}
	tp := New(mock, Makecho())

	tp.addtype([]byte("h"), 0, true)
	tp.Sendtype()

	testhelpers.AssertInt(t, 0, len(mock.writes))
	testhelpers.AssertString(t, "h", string(tp.Typing))
}

func Test_SendtypeMultiblock(t *testing.T) {
	mock := &mockttyfd{make([][]byte, 0, 10)}
	tp := New(mock, Makecho())

	tp.addtype([]byte("hello\nworld\nbye"), 0, true)
	tp.Sendtype()

	testhelpers.AssertInt(t, 2, len(mock.writes))
	testhelpers.AssertString(t, "hello\n", string(mock.writes[0]))
	testhelpers.AssertString(t, "world\n", string(mock.writes[1]))
	testhelpers.AssertString(t, "bye", string(tp.Typing))
}

func Test_Type(t *testing.T) {
	mock := &mockttyfd{make([][]byte, 0, 10)}
	tp := New(mock, Makecho())

	e := &acme.Event{Nr: len("hello"), Text: []byte("hello")}
	tp.Type(e)

	testhelpers.AssertString(t, "hello", string(tp.Typing))
}

func Test_TypeCook(t *testing.T) {
	mock := &mockttyfd{make([][]byte, 0, 10)}
	tp := New(mock, Makecho())

	s := "hello\n"
	e := &acme.Event{Nr: len(s), Text: []byte(s)}
	tp.Type(e)

	testhelpers.AssertInt(t, 1, len(mock.writes))
	testhelpers.AssertString(t, "hello\n", string(mock.writes[0]))
	testhelpers.AssertString(t, "", string(tp.Typing))
	testhelpers.AssertBool(t, true, tp.cook)
}
