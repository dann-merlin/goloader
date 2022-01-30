//go:build go1.16 && !go1.18
// +build go1.16,!go1.18

package goloader

// A funcID identifies particular functions that need to be treated
// specially by the runtime.
// Note that in some situations involving plugins, there may be multiple
// copies of a particular special runtime function.
// Note: this list must match the list in cmd/internal/objabi/funcid.go.
type funcID uint8

// Layout of in-memory per-function information prepared by linker
// See https://golang.org/s/go12symtab.
// Keep in sync with linker (../cmd/link/internal/ld/pcln.go:/pclntab)
// and with package debug/gosym and with symtab.go in package runtime.
type _func struct {
	entry   uintptr // start pc
	nameoff int32   // function name

	args        int32  // in/out args size
	deferreturn uint32 // offset of start of a deferreturn call instruction from entry, if any.

	pcsp      uint32
	pcfile    uint32
	pcln      uint32
	npcdata   uint32
	cuOffset  uint32  // runtime.cutab offset of this function's CU
	funcID    funcID  // set for certain special runtime functions
	_         [2]byte // pad
	nfuncdata uint8   // must be last
}

func initfunc(symbol *ObjSymbol, nameOff, spOff, pcfileOff, pclnOff, cuOff int) _func {
	fdata := _func{
		entry:       uintptr(0),
		nameoff:     int32(nameOff),
		args:        int32(symbol.Func.Args),
		deferreturn: uint32(0),
		pcsp:        uint32(spOff),
		pcfile:      uint32(pcfileOff),
		pcln:        uint32(pclnOff),
		npcdata:     uint32(len(symbol.Func.PCData)),
		cuOffset:    uint32(cuOff),
		funcID:      funcID(symbol.Func.FuncID),
		nfuncdata:   uint8(len(symbol.Func.FuncData)),
	}
	return fdata
}

func setfuncentry(f *_func, entry uintptr, text uintptr) {
	f.entry = entry
}
