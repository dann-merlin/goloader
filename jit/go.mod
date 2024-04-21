module github.com/dann-merlin/goloader/jit

go 1.18

require (
	github.com/dann-merlin/goloader v0.0.0-20240111193324-64f971021f52
	github.com/dann-merlin/goloader/jit/testdata v0.0.0-20240111193324-64f971021f52
	github.com/dann-merlin/goloader/unload v0.0.0-20240111193324-64f971021f52
)

require github.com/opentracing/opentracing-go v1.2.0 // indirect

//replace github.com/dann-merlin/goloader => ../
//replace github.com/dann-merlin/goloader/jit/testdata => ./testdata
replace github.com/dann-merlin/goloader/jit/objfile => ./objfile
