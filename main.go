package main

import (
	"github.com/FruitSpace/GhostCore.Modules/bundler"
	"github.com/FruitSpace/GhostCore.Modules/engine"
	esbuild "github.com/evanw/esbuild/pkg/api"
	"log"
	"os"
)

var scripts = []string{
	"test_function.js",
	"test_anonymous.js",
	"test_modules_compiled.js",
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
	fname := os.Args[1]
	data, err := bundler.AutoFolder(fname)
	if err != nil {
		log.Panicln(err)
	}
	log.Printf("Bundling '%s' (ver.%s) by %s", data.Script, data.Version, data.Author)
	nb := bundler.NewBundler().WithEntryPoint(data.Entrypoint).UseMinify(true).
		WithOutfile("script.js", false).WithLogLevel(esbuild.LogLevelInfo)
	code, err := nb.Bundle()
	if err != nil {
		log.Println(err)
		log.Println(nb.GetErrors())
		return
	}
	log.Println("Running in VM...")
	e := engine.NewEngine()
	err = e.LoadString(string(code.Contents))
	if err != nil {
		log.Panicln(err)
	}
	for {
	}
}
