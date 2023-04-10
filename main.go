package main

import (
	"github.com/FruitSpace/GhostCore.Modules/goja_engine"
	"github.com/dop251/goja"
	"log"
	"reflect"
)

var scripts = []string{
	"test_function.js",
	"test_anonymous.js",
	"test_modules.js",
	//"test_math.js",
}

var testF = map[string]interface{}{
	"fNumber": 40,
	"fString": "Hello",
	"fObject": struct {
		Data string
	}{"This is object data"},
	"fFunction": "deferred value yay",
}

func main() {
	for _, script := range scripts {
		log.Println("Testing", script)
		engine := goja_engine.NewEngine()
		err := engine.LoadFile("./js_modules/" + script)
		if err != nil {
			log.Panicln(err)
		}
		for k, v := range testF {
			ret, err := engine.Call(k, v)
			if err != nil {
				log.Panicln(err)
			}
			log.Printf("%s(%v) -> %+v", k, v, ret)
			if reflect.TypeOf(ret) == reflect.TypeOf((func(goja.FunctionCall) goja.Value)(nil)) {
				log.Printf("╰ Reflect detected function. Output: %+v",
					ret.(func(goja.FunctionCall) goja.Value)(goja.FunctionCall{}).Export())
			}
		}
	}
}
