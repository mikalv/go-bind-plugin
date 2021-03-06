package cli

// Autogenerated by github.com/wendigo/go-bind-plugin on 2016-11-10 13:43:16.162648817 +0100 CET, do not edit!
// Command: go-bind-plugin -plugin-path plugin.so -plugin-package ../internal/test_fixtures/benchmark_plugin -output-name BenchmarkPlugin -output-path cli_test_plugin.go -output-package cli
//
// Plugin plugin.so info:
// - package: github.com/wendigo/go-bind-plugin/internal/test_fixtures/benchmark_plugin
// - size: 1230240 bytes
// - sha256: a97d2f90b6c636f7efb93d677e679c143f4a508af4d1fa49c5c91a1467da7e0d

import (
	"fmt"

	"strings"

	"plugin"

	"reflect"
)

// BenchmarkPlugin wraps symbols (functions and variables) exported by plugin github.com/wendigo/go-bind-plugin/internal/test_fixtures/benchmark_plugin
//
// See docs at https://godoc.org/github.com/wendigo/go-bind-plugin/internal/test_fixtures/benchmark_plugin
type BenchmarkPlugin struct {
	// Exported functions
	_Sum func([]uint64) uint64

	// Exported variables (public references)

}

// Sum function was exported from plugin github.com/wendigo/go-bind-plugin/internal/test_fixtures/benchmark_plugin symbol 'Sum'
//
// See docs at https://godoc.org/github.com/wendigo/go-bind-plugin/internal/test_fixtures/benchmark_plugin#Sum
func (p *BenchmarkPlugin) Sum(in0 []uint64) uint64 {
	return p._Sum(in0)
}

// String returnes textual representation of the wrapper. It provides info on exported symbols and variables.
func (p *BenchmarkPlugin) String() string {
	var lines []string
	lines = append(lines, "Struct BenchmarkPlugin:")
	lines = append(lines, "\t- Generated on: 2016-11-10 13:43:16.162648817 +0100 CET")
	lines = append(lines, "\t- Command: go-bind-plugin -plugin-path plugin.so -plugin-package ../internal/test_fixtures/benchmark_plugin -output-name BenchmarkPlugin -output-path cli_test_plugin.go -output-package cli")
	lines = append(lines, "\nPlugin info:")
	lines = append(lines, "\t- package: github.com/wendigo/go-bind-plugin/internal/test_fixtures/benchmark_plugin")
	lines = append(lines, "\t- sha256 sum: a97d2f90b6c636f7efb93d677e679c143f4a508af4d1fa49c5c91a1467da7e0d")
	lines = append(lines, "\t- size: 1230240 bytes")
	lines = append(lines, "\nExported functions (1):")
	lines = append(lines, "\t- Sum func([]uint64) (uint64)")

	return strings.Join(lines, "\n")
}

// BindBenchmarkPlugin loads plugin from the given path and binds symbols (variables and functions)
// to the BenchmarkPlugin struct.
func BindBenchmarkPlugin(path string) (*BenchmarkPlugin, error) {
	p, err := plugin.Open(path)

	if err != nil {
		return nil, fmt.Errorf("Could not open plugin: %s", err)
	}

	ret := new(BenchmarkPlugin)

	funcSum, err := p.Lookup("Sum")
	if err != nil {
		return nil, fmt.Errorf("Could not import function 'Sum', symbol not found: %s", err)
	}

	if typed, ok := funcSum.(func([]uint64) uint64); ok {
		ret._Sum = typed
	} else {
		return nil, fmt.Errorf("Could not import function 'Sum', incompatible types 'func([]uint64) (uint64)' and '%s'", reflect.TypeOf(funcSum))
	}

	return ret, nil
}
