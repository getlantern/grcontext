// Package grcontext provides a mechanism for binding a context.Context
// to the current goroutine.
package grcontext

import (
	"context"
	"sync"
)

var (
	contexts   = make(map[uint64]context.Context)
	mxContexts sync.RWMutex
)

// Bind binds the given context to the current goroutine and returns a function
// that unbinds it. If a context was already bound to this goroutine, when the
// the new context is unbound, the prior context is automatically rebound. So,
// it's possible to nest multiple calls to Bind on the same goroutine.
func Bind(ctx context.Context) func() {
	id := curGoroutineID()
	mxContexts.Lock()
	defer mxContexts.Unlock()
	prior, hasPrior := contexts[id]
	contexts[id] = ctx

	return func() {
		mxContexts.Lock()
		defer mxContexts.Unlock()
		if !hasPrior {
			delete(contexts, id)
		} else {
			contexts[id] = prior
		}
	}
}

// Current gets the context bound to the current goroutine, or context.TODO()
// if none is bound.
func Current() context.Context {
	id := curGoroutineID()
	mxContexts.RLock()
	defer mxContexts.RUnlock()
	current, hasCurrent := contexts[id]
	if !hasCurrent {
		return context.TODO()
	}
	return current
}
