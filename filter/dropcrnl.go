// Assorted filters
package filter

func Dropcrnl(r []byte) []byte {
	w := make([]byte, 0, len(r))
	
	for i := 0 ; i < len(r); i++ {
		if i +1 < len(r) && r[i] == '\r' && r[i+1] == '\n' {
			i ++
			continue
		} else {
			w = append(w, r[i])
		}
	}
	return w
}

