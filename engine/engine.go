package engine

import (
	"errors"
	"fmt"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/eventloop"
	"github.com/dop251/goja_nodejs/require"
	babel "github.com/jvatic/goja-babel"
	"github.com/olebedev/gojax/fetch"
	"gopkg.in/elazarl/goproxy.v1"
	"log"
	"os"
)

type Engine struct {
	vm   *goja.Runtime
	loop *eventloop.EventLoop
}

func NewEngine() *Engine {
	loop := eventloop.NewEventLoop()
	loop.Start()
	vm := goja.New()
	new(require.Registry).Enable(vm)
	console.Enable(vm)
	fetch.Enable(loop, goproxy.NewProxyHttpServer())
	babel.Init(4)
	//registry := new(require.Registry)
	//_ = registry.Enable(vm)
	return &Engine{
		vm:   vm,
		loop: loop,
	}
}

func (e *Engine) Call(name string, arg interface{}) (interface{}, error) {
	v := e.vm.Get(name)
	f, ok := goja.AssertFunction(v)
	if !ok {
		return nil, errors.New(fmt.Sprintf("%s(%v) is not a function", name, arg))
	}
	ret, err := f(goja.Undefined(), e.vm.ToValue(arg))
	return ret.Export(), err
}

func (e *Engine) GetVM() *goja.Runtime {
	return e.vm
}

func (e *Engine) LoadFile(path string) error {
	var code string
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	code = string(data)
	//code, err = e.TranspileBabel(code)
	//if err != nil {
	//	return err
	//}
	e.loop.RunOnLoop(func(vm *goja.Runtime) {
		a, _ := vm.RunString(code)
		log.Printf("Load %s: %+v\n", path, a)
	})

	return err
}

func (e *Engine) LoadString(code string) error {
	//code, err := e.TranspileBabel(code)
	//if err != nil {
	//	return err
	//}
	fmt.Println(code)
	e.loop.RunOnLoop(func(vm *goja.Runtime) {
		a, err := vm.RunString(code)
		log.Printf("Load string: %+v\n", a)
		if err != nil {
			log.Panicln(err)
		}
	})
	return nil
}

func (e *Engine) TranspileBabel(code string) (string, error) {
	return babel.TransformString(code, map[string]interface{}{
		"presets": []string{
			"env",
		},
	})
}
