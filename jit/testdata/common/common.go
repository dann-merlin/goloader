package common

import (
	"github.com/opentracing/opentracing-go"
	"sync"
)

var Val = 99

func Inc() {
	Val += 1
}

func Dec() {
	Val += 1
}

type Pointer *struct{}

type SomeStruct struct {
	Val1    interface{}
	Val2    map[string]interface{}
	Val3    []error
	Val4    opentracing.Span
	Mutex   *sync.Mutex
	Pointer Pointer
}

type SomeInterface interface {
	Method1(input SomeStruct) (SomeStruct, error)
	Method2(input map[string]interface{}) error
}

type StartStoppable interface {
	Start() error
	Stop() error
	InChan() chan<- SomeStruct
	OutChan() <-chan SomeStruct
}

type MessageWriter interface {
	Dial(addr string) error
	Close() error
	WriteMessage(data string) (int, error)
}
