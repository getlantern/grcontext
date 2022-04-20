# grcontext
Adds support for binding context.Contexts's to the current goroutine.

This library is inspired by https://github.com/tylerb/gls and https://github.com/jtolds/gls.

## Example

```go
ctx := context.WithValue(context.Background(), "key", "value")

unbind := grcontext.Bind(ctx)
if grcontext.Current() != ctx {
    panic("current context should be ctx")
}

unbind()
if grcontext.Current() != context.TODO() {
    panic("current context should be context.TODO()")
}
```