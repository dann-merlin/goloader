//go:build go1.8 && !go1.14
// +build go1.8,!go1.14

package goloader

func (linker *Linker) addDeferReturn(_func *_func) (err error) {
	return nil
}
