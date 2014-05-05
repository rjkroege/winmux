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

// Positions 0, 1 are before the buffer.
// Position 2 is the end of the buffer.
func Test_Addtyping(t *testing.T) {
	ws := New()
	ws.Move(2)

	testhelpers.AssertBool(t, true, ws.Beforeslice(0))
	testhelpers.AssertBool(t, true, ws.Beforeslice(1))

	ws.Addtyping([]byte{'a'}, 2)
	testhelpers.AssertString(t, "a", ws.String())

	testhelpers.AssertBool(t, true, ws.Beforeslice(1))

	testhelpers.AssertBool(t, true, ws.Inslice(2))
	testhelpers.AssertBool(t, true, ws.Inslice(3))
	testhelpers.AssertBool(t, false, ws.Inslice(4))

	testhelpers.AssertInt(t, 2, ws.Offset)
	testhelpers.AssertBool(t, true, ws.Afterslice(2, 0))
	testhelpers.AssertBool(t, false, ws.Afterslice(1, 1))

	ws.Addtyping([]byte{'b'}, 3)
	testhelpers.AssertString(t, "ab", ws.String())
	testhelpers.AssertBool(t, true, ws.Inslice(4))
	testhelpers.AssertBool(t, false, ws.Inslice(5))

	p, q := ws.Extent()
	testhelpers.AssertInt(t, 2, p)
	testhelpers.AssertInt(t, 4, q)

	ws.Addtyping([]byte{'c'}, 4)
	testhelpers.AssertString(t, "abc", ws.String())
	testhelpers.AssertInt(t, len("abc"), ws.Ntyper())

	testhelpers.AssertInt(t, 2, ws.Offset)
	ws.Addtyping([]byte{'X'}, 2)
	testhelpers.AssertString(t, "Xabc", ws.String())
	testhelpers.AssertInt(t, len("Xabc"), ws.Ntyper())

	ws.Addtyping([]byte{'Y'}, 3)
	testhelpers.AssertString(t, "XYabc", ws.String())

}

func Test_Reset(t *testing.T) {
	ws := New()
	ws.Move(2)
	ws.Addtyping([]byte{'b'}, 2)

	ws.Reset()
	testhelpers.AssertString(t, "", ws.String())
}

func Test_Addtyping_BeforePanic(t *testing.T) {
	ws := New()
	ws.Move(2)

	defer func() {
		if e := recover(); e != nil {
			s := e.(string)
			testhelpers.AssertString(t, "p (0) !in [ws.Offset: 2, ws.Offset + len: 2)", s)
		} else {
			t.Fail()
		}
	}()

	ws.Addtyping([]byte{'c'}, 0)
}

func Test_Addtyping_AfterPanic(t *testing.T) {
	ws := New()
	ws.Move(2)

	defer func() {
		if e := recover(); e != nil {
			s := e.(string)
			testhelpers.AssertString(t, "p (3) !in [ws.Offset: 2, ws.Offset + len: 2)", s)
		} else {
			t.Fail()
		}
	}()

	ws.Addtyping([]byte{'c'}, 3)
}


func Test_Delete_to_empty(t *testing.T) {
	ws := New()
	ws.Move(2)
	ws.Addtyping([]byte{'a'}, 2)

	n := ws.Delete(2, 3)
	testhelpers.AssertString(t, "", ws.String())
	testhelpers.AssertInt(t, 0, n)
}

func Test_Delete_in_middle(t *testing.T) {
	ws := New()
	ws.Move(2)
	ws.Addtyping([]byte{'a', 'b', 'c'}, 2)
	testhelpers.AssertString(t, "abc", ws.String())

	n := ws.Delete(3,4)
	testhelpers.AssertString(t, "ac", ws.String())
	testhelpers.AssertInt(t, 0, n)
}

func Test_Delete_before_offset(t *testing.T) {
	ws := New()
	ws.Move(2)

	n := ws.Delete(1,2)
	testhelpers.AssertString(t, "", ws.String())
	testhelpers.AssertInt(t, 1, n)
}

func Test_Delete_spanning_offset(t *testing.T) {
	ws := New()
	ws.Move(2)
	ws.Addtyping([]byte{'a', 'b', 'c'}, 2)
	testhelpers.AssertString(t, "abc", ws.String())

	n := ws.Delete(1,3)
	testhelpers.AssertString(t, "bc", ws.String())
	testhelpers.AssertInt(t, 1, n)
}


func Test_Delete_multi(t *testing.T) {
	ws := New()
	ws.Move(2)
	ws.Addtyping([]byte{'a', 'b', 'c'}, 2)
	testhelpers.AssertString(t, "abc", ws.String())

	n := ws.Delete(3,5)
	testhelpers.AssertString(t, "a", ws.String())
	testhelpers.AssertInt(t, 0, n)
}


func Test_Delete_panic(t *testing.T) {
	ws := New()
	ws.Move(2)
	ws.Addtyping([]byte{'a', 'b', 'c'}, 2)
	testhelpers.AssertString(t, "abc", ws.String())

	defer func() {
		if e := recover(); e != nil {
			s := e.(string)
			testhelpers.AssertString(t, "Delete went wrong", s)
		} else {
			t.Fail()
		}
	}()

	ws.Delete(5,6)
}
