package engine

import (
	"github.com/dop251/goja"
	"sync"
)

var once = &sync.Once{}
var runtimepool chan *Engine
var scriptcache map[string]*goja.Program

func InitRuntimePool(size int) (err error) {
	defer (func() {
		if err != nil {
			once = &sync.Once{}
		}
	})()
	once.Do(func() {

	})

	return err
}
