package tendo

type root struct {
	currentLibrary string
	libraries      map[string]*library
}

func newRoot() *root {
	return &root{
		libraries: make(map[string]*library),
	}
}
