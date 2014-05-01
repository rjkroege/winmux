// Assorted filters
package filter

// TODO(rjkroege): needs to do the right thing with
// bs characters too.
func Dropcrnl(r []byte) []byte {
	w := make([]byte, 0, len(r))
	
	for i := 0 ; i < len(r); i++ {
		if i +1 < len(r) && r[i] == '\r' && r[i+1] == '\n' {
			continue
		}
		w = append(w, r[i])
	}
	return w
}

