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
