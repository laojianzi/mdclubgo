package goroutine

import (
	"github.com/laojianzi/mdclubgo/middleware"
)

// Go runs the given function in a goroutine and catches and logs panics.
//
// This prevents a single panicking goroutine from crashing the entire binary
func Go(f func()) {
	go func() {
		defer middleware.RecoverHandle(nil)

		f()
	}()
}
