// +build go1.16
// +build !go1.17

package goloader

import (
	"strings"
	"unsafe"
)

func registerFunc(md *moduledata, symPtr map[string]uintptr) {
	//register function
	for _, f := range md.ftab {
		if int(f.funcoff) < len(md.pclntable) {
			_func := (*_func)(unsafe.Pointer((&(md.pclntable[f.funcoff]))))
			name := gostringnocopy(&(md.funcnametab[_func.nameoff]))
			if !strings.HasPrefix(name, TypeDoubleDotPrefix) && _func.entry < md.etext {
				symPtr[name] = _func.entry
			}
		}
	}
}
