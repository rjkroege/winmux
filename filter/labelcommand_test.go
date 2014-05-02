package filter

import (
	"github.com/rjkroege/wikitools/testhelpers"
	"testing"
)

func Test_Labelcommand_basic(t *testing.T) {
	b, r := Labelcommand([]byte{})
	if r != nil {
		t.Error("empty buf should have no label")
	}
	if len(b) != 0 {
		t.Error("empty buf should still be empty")
	}

	b, r = Labelcommand([]byte("h"))
	if r != nil {
		t.Error("simple buf should have no label")
	}
	testhelpers.AssertString(t, "h", string(b))
}

func Test_Labelcommand_hasend(t *testing.T) {
	b, r := Labelcommand([]byte("he\007llo"))
	if r != nil {
		t.Error("incomplete command should have no label")
	}
	testhelpers.AssertString(t, "he\007llo", string(b))
}

func Test_Labelcommand_hasend_with_space(t *testing.T) {
	b, r := Labelcommand([]byte("hel\007lo"))
	if r != nil {
		t.Error("incomplete command should have no label")
	}
	testhelpers.AssertString(t, "hel\007lo", string(b))
}

func Test_Labelcommand_has_empty(t *testing.T) {
	b, r := Labelcommand([]byte("he\033];\007llo"))
	testhelpers.AssertString(t, "hello", string(b))
	testhelpers.AssertString(t, "", string(r))
}

func Test_Labelcommand_hascmd(t *testing.T) {
	b, r := Labelcommand([]byte("he\033];world\007llo"))
	testhelpers.AssertString(t, "hello", string(b))
	testhelpers.AssertString(t, "world", string(r))
}

func Test_Labelcommand_hascmd_start(t *testing.T) {
	b, r := Labelcommand([]byte("\033];world\007hello"))
	testhelpers.AssertString(t, "hello", string(b))
	testhelpers.AssertString(t, "world", string(r))
}

func Test_Labelcommand_hascmd_end(t *testing.T) {
	b, r := Labelcommand([]byte("hello\033];world\007"))
	testhelpers.AssertString(t, "hello", string(b))
	testhelpers.AssertString(t, "world", string(r))
}
