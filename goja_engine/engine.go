package goja_engine

import (
	"errors"
	"fmt"
	"github.com/dop251/goja"
	babel "github.com/jvatic/goja-babel"
	"log"
	"os"
)

type Engine struct {
	vm *goja.Runtime
}

func NewEngine() *Engine {
	vm := goja.New()
	babel.Init(4)
	//registry := new(require.Registry)
	//_ = registry.Enable(vm)
	return &Engine{
		vm: vm,
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

func (e *Engine) LoadFile(path string) error {
	var code string
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	code = string(data)
	code, err = e.TranspileBabel(code)
	if err != nil {
		return err
	}
	a, err := e.vm.RunString(code)
	log.Printf("Load %s: %+v\n", path, a)
	return err
}

func (e *Engine) TranspileBabel(code string) (string, error) {
	return babel.TransformString(code, map[string]interface{}{
		"presets": []string{
			"env",
		},
	})
}
