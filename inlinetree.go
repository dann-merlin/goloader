// +build go1.9
// +build !go1.15

package goloader

import (
	"cmd/objfile/goobj"
	"unsafe"
)

func findFileTab(codereloc *CodeReloc, filename string) int32 {
	tab := codereloc.namemap[filename]
	for index, value := range codereloc.filetab {
		if uint32(tab) == value {
			return int32(index)
		}
	}
	return -1
}

func _addInlineTree(codereloc *CodeReloc, _func *_func, symbol *goobj.Sym, fd *readAtSeeker) (err error) {
	funcname := symbol.Name
	Func := symbol.Func
	sym := codereloc.symMap[funcname]
	if Func != nil && len(Func.InlTree) != 0 {
		name := funcname + INLINETREE_SUFFIX

		for _func.npcdata <= _PCDATA_InlTreeIndex {
			sym.Func.PCData = append(sym.Func.PCData, uint32(0))
			_func.npcdata++
		}
		sym.Func.PCData[_PCDATA_InlTreeIndex] = uint32(len(codereloc.pclntable))

		fd.ReadAtWithSize(&(codereloc.pclntable), symbol.Func.PCInline.Size, symbol.Func.PCInline.Offset)
		for _, inl := range symbol.Func.InlTree {
			if _, ok := codereloc.namemap[inl.Func.Name]; !ok {
				codereloc.namemap[inl.Func.Name] = len(codereloc.pclntable)
				codereloc.pclntable = append(codereloc.pclntable, []byte(inl.Func.Name)...)
				codereloc.pclntable = append(codereloc.pclntable, ZERO_BYTE)
			}
		}

		bytes := make([]byte, len(Func.InlTree)*InlinedCallSize)
		for k, inl := range Func.InlTree {
			inlinedcall := initInlinedCall(codereloc, inl, _func)
			copy2Slice(bytes[k*InlinedCallSize:], uintptr(unsafe.Pointer(&inlinedcall)), InlinedCallSize)
		}
		codereloc.stkmaps[name] = bytes
		for _func.nfuncdata <= _FUNCDATA_InlTree {
			sym.Func.FuncData = append(sym.Func.FuncData, uintptr(0))
			_func.nfuncdata++
		}
		sym.Func.FuncData[_FUNCDATA_InlTree] = (uintptr)(unsafe.Pointer(&(codereloc.stkmaps[name][0])))
	}
	return err
}
