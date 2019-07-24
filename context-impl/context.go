package contextimpl

import (
	"errors"
	"reflect"
	"sync"
	"time"
)

type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <- chan struct{}
	Err() error
	Value(key interface{}) interface{}
}

type emptyCtx int

func (emptyCtx) Deadline() (deadline time.Time, ok bool) { return }
func (emptyCtx) Done() <- chan struct{} { return nil }
func (emptyCtx) Err() error { return nil }
func (emptyCtx) Value(key interface{}) interface{} { return nil }

var (
	backgroud = new(emptyCtx)
	todo = new(emptyCtx)
)

func Background() Context {
	return backgroud
}

func TODO() Context {
	return todo
}

type cancelCtx struct {
	Context
	done chan struct{}
	err error
	mu sync.Mutex
}

func (ctx *cancelCtx) Done() <- chan struct{} { return ctx.done }
func (ctx *cancelCtx) Err() error {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	return ctx.err
}

var Canceled = errors.New("context canceled")

type CancelFunc func()

func WithCancel(parent Context) (Context, CancelFunc) {
	ctx := &cancelCtx{
		Context: parent,
		done: make(chan struct{}),
	}

	cancel := func() { ctx.cancel(Canceled) }

	go func() {
		select {
		case <-parent.Done():
			ctx.cancel(parent.Err())
		case <- ctx.Done():
		}
	}()

	return ctx, cancel
}

func (ctx *cancelCtx) cancel(err error) {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()

	if ctx.err != nil {
		return
	}
	ctx.err = err
	close(ctx.done)
}

type deadlineCtx struct {
	*cancelCtx
	deadline time.Time
}

func (ctx *deadlineCtx) Deadline() (deadline time.Time, ok bool) {
	return ctx.deadline, true
}

var DeadlineExceeded = errors.New("deadline exceeded")

func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc) {
	cctx, cancel := WithCancel(parent)

	ctx := &deadlineCtx{
		cancelCtx: cctx.(*cancelCtx),
		deadline: deadline,
	}

	t := time.AfterFunc(time.Until(deadline), func() {
		ctx.cancel(DeadlineExceeded)
	})

	stop := func() {
		t.Stop()
		cancel()
	}

	return ctx, stop
}

func WithTimeout(parent Context, duration time.Duration) (Context, CancelFunc) {
	return WithDeadline(parent, time.Now().Add(duration))
}

type valueCtx struct {
	Context
	value, key interface{}
}

func (ctx *valueCtx) Value(key interface{}) interface{} {
	if key == ctx.key {
		return ctx.value
	}

	return ctx.Context.Value(key)
}

func WithValue(parent Context, key, value interface{}) Context {
	if key == nil {
		panic("key is nil")
	}
	if !reflect.TypeOf(key).Comparable() {
		panic("key is not comparable")
	}

	return &valueCtx{
		Context: parent,
		value: value,
		key: key,
	}
}
