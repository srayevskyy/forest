package forest

import (
	"fmt"
	"os"
	"strings"
)

// T is the interface that this package is using from standard testing.T
type T interface {
	// Logf formats its arguments according to the format, analogous to Printf, and records the text in the error log.
	// The text will be printed only if the test fails or the -test.v flag is set.
	Logf(format string, args ...interface{})
	// Error is equivalent to Log followed by Fail.
	Error(args ...interface{})
	// Fatal is equivalent to Log followed by FailNow.
	Fatal(args ...interface{})
}

// TestingT provides a sub-api of testing.T
var TestingT = logging{true}

type logging struct {
	// this field exists for testing this package only
	doExit bool
}

func (f logging) Logf(format string, args ...interface{}) {
	fmt.Printf("\tinfo : "+tabify(format)+"\n", args...)
}

func (f logging) Error(args ...interface{}) {
	fmt.Printf("\terror: "+tabify("%s")+"\n", args)
}

func (f logging) Fatal(args ...interface{}) {
	fmt.Printf("\tfatal: "+tabify("%s")+"\n", args...)
	if f.doExit {
		os.Exit(1)
	}
}

func tabify(format string) string {
	if strings.HasPrefix(format, "\n") {
		return strings.Replace(format, "\n", "\n\t\t", 1)
	}
	return format
}
