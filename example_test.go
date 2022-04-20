package grcontext_test

import (
	"context"

	"github.com/getlantern/grcontext"
)

func ExampleBind() {
	ctx := context.WithValue(context.Background(), "key", "value")

	unbind := grcontext.Bind(ctx)
	if grcontext.Current() != ctx {
		panic("current context should be ctx")
	}

	unbind()
	if grcontext.Current() != context.TODO() {
		panic("current context should be context.TODO()")
	}
}
