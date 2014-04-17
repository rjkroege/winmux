// Assorted filters
package filter

func Squashnul(r []byte) []byte {
	w := make([]byte, 0, len(r))
	
	for i := 0 ; i < len(r); i++ {
		if r[i] == '\000' {
			continue
		} else {
			w = append(w, r[i])
		}
	}
	return w
}

