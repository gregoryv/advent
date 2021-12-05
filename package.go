package advent

import "log"

func shouldNot(err error) {
	if err != nil {
		fatalFunc(err)
	}
}

// Used by shouldNot
var fatalFunc func(...interface{}) = log.Fatal

// SetDebug(nil) to reset
func SetDebug(v logPrinter) {
	if v == nil {
		debugOn = false
		debug = &discard{}
		return
	}
	debugOn = true
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
