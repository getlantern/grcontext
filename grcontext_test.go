package grcontext

import (
	"context"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBind(t *testing.T) {
	ctx1 := context.WithValue(context.Background(), "key1", "value1")
	ctx2 := context.WithValue(context.Background(), "key2", "value2")
	ctx3 := context.WithValue(context.Background(), "key3", "value3")

	unbind1 := Bind(ctx1)
	require.Equal(t, ctx1, Current(), "should bind correct context")
	unbind2 := Bind(ctx2)
	require.Equal(t, ctx2, Current(), "should override already bound context")

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		unbind3 := Bind(ctx3)
		require.Equal(t, ctx3, Current(), "should bind correct context")
		unbind3()
		require.Equal(t, context.TODO(), Current(), "should not see context bound to other goroutine")
	}()
	wg.Wait()

	require.Equal(t, ctx2, Current(), "should be unaffected by contexts bound in other goroutine")
	unbind2()
	require.Equal(t, ctx1, Current(), "should revert to previously bound context")
	unbind1()
	require.Equal(t, context.TODO(), Current(), "should have no bound context")
}
