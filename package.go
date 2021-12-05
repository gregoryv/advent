package advent

import "log"

func shouldNot(err error) {
	if err == nil {
		return
	}
	fatalHandler(err)
}

// Used by shouldNot
var fatalHandler func(...interface{}) = log.Fatal

func SetDebug(v logPrinter) {
	debugOn = v == nil
	debug = v
}

var debugOn bool

// useful to debug things during testing, replace it with testing.T
var debug logPrinter = &discard{}

type discard struct{}

func (*discard) Log(...interface{})          {}
func (*discard) Logf(string, ...interface{}) {}

type logPrinter interface {
	Log(v ...interface{})
	Logf(format string, v ...interface{})
}
