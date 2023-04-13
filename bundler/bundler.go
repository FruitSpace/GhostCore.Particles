package bundler

import (
	"errors"
	"fmt"
	esbuild "github.com/evanw/esbuild/pkg/api"
	"log"
)

func NewBundler() *Bundler {
	return &Bundler{
		outfile:   "dist/gcmodule.js",
		minify:    true,
		writeDisk: false,
		logLevel:  esbuild.LogLevelError,
	}
}

type Bundler struct {
	outfile     string
	minify      bool
	minifyNames bool
	EntryPoints []string
	writeDisk   bool
	logLevel    esbuild.LogLevel
	errors      []esbuild.Message
}

type OutputFile struct {
	esbuild.OutputFile
}

func (b *Bundler) WithOutfile(outfile string, toDisk bool) *Bundler {
	b.writeDisk = toDisk
	b.outfile = outfile
	return b
}

func (b *Bundler) WithLogLevel(level esbuild.LogLevel) *Bundler {
	b.logLevel = level
	return b
}

func (b *Bundler) WithEntryPoints(entrypoints []string) *Bundler {
	b.EntryPoints = make([]string, len(entrypoints))
	copy(b.EntryPoints, entrypoints)
	return b
}

func (b *Bundler) WithEntryPoint(entrypoint string) *Bundler {
	return b.WithEntryPoints([]string{entrypoint})
}

func (b *Bundler) UseMinify(minify bool) *Bundler {
	b.minify = minify
	return b
}

func (b *Bundler) bundle() esbuild.BuildResult {
	b.errors = make([]esbuild.Message, 0)
	return esbuild.Build(esbuild.BuildOptions{
		Outfile:           b.outfile,
		MinifyIdentifiers: b.minifyNames,
		MinifySyntax:      b.minify,
		MinifyWhitespace:  b.minify,
		EntryPoints:       b.EntryPoints,
		Bundle:            true,
		LogLevel:          b.logLevel,
	})
}

func (b *Bundler) Bundle() (OutputFile, error) {
	result := b.bundle()
	b.errors = result.Errors
	if len(b.errors) > 0 {
		err := b.errors[0]
		return OutputFile{}, errors.New(fmt.Sprintf("[%s] An error occured:\n%s", err.PluginName, err.Text))
	}
	if len(result.OutputFiles) != 1 {
		log.Printf("[Anomaly] Output file count = %d\n", len(result.OutputFiles))
		for _, f := range result.OutputFiles {
			log.Println("-->", f.Path)
		}
	}

	return OutputFile{result.OutputFiles[0]}, nil
}

func (b *Bundler) GetErrors() []esbuild.Message {
	return b.errors
}
