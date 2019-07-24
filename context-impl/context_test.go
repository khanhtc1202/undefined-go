package contextimpl

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestBackgroundNotTodo(t *testing.T) {
	todo := fmt.Sprint(TODO())
	bg := fmt.Sprint(Background())

	if todo == bg {
		t.Errorf("TODO and BG are equal: %q vs %q", todo, bg)
	}
}

func TestWithCancel(t *testing.T) {
	ctx, cancel := WithCancel(Background())

	if err := ctx.Err(); err != nil {
		t.Errorf("error should be nil, got %v", err)
	}
	cancel()
	<- ctx.Done()
	if err := ctx.Err(); err != Canceled {
		t.Errorf("error should be canceled now, got %v", err)
	}
}

func TestWithCancelConcurrent(t *testing.T) {
	ctx, cancel := WithCancel(Background())

	time.AfterFunc(1*time.Second, cancel)

	if err := ctx.Err(); err != nil {
		t.Errorf("error should be nil first, got %v", err)
	}
	<- ctx.Done()
	if err := ctx.Err(); err != Canceled {
		t.Errorf("error should be canceled now, got %v", err)
	}
}

func TestWithCancelPropagation(t *testing.T) {
	ctxA, cancelA := WithCancel(Background())
	ctxB, cancelB := WithCancel(ctxA)
	defer cancelB()

	cancelA()
	select {
	case <-ctxB.Done():
	case <-time.After(1*time.Second):
		t.Errorf("time out")
	}

	if err := ctxB.Err(); err != Canceled {
		t.Errorf("error should be canceled, got %v", err)
	}
}

func TestWithDeadline(t *testing.T) {
	deadline := time.Now().Add(2 * time.Second)
	ctx, cancel := WithDeadline(Background(), deadline)

	if d, ok := ctx.Deadline(); !ok || d != deadline {
		t.Errorf("expect deadline %v, got %v", deadline, d)
	}

	then := time.Now()
	<- ctx.Done()
	if d := time.Since(then); math.Abs(d.Seconds()-2.0) > 0.1 {
		t.Errorf("should have been done after 2.0 seconds, took %v", d)
	}

	if err := ctx.Err(); err != DeadlineExceeded {
		t.Errorf("error should be DeadlineExceeded, got %v", err)
	}

	cancel()
	if err := ctx.Err(); err != DeadlineExceeded {
		t.Errorf("error should still be DeadlineExceeded, got %v", err)
	}
}

func TestWithTimeout(t *testing.T) {
	duration := time.Duration(2 * time.Second)
	deadline := time.Now().Add(duration)
	ctx, cancel := WithTimeout(Background(), duration)

	if d, ok := ctx.Deadline(); !ok || d.Sub(deadline) > time.Microsecond {
		t.Errorf("expect deadline %v, got %v", deadline, d)
	}

	then := time.Now()
	<- ctx.Done()
	if d := time.Since(then); math.Abs(d.Seconds()-2.0) > 0.1 {
		t.Errorf("should have been done after 2.0 seconds, took %v", d)
	}

	if err := ctx.Err(); err != DeadlineExceeded {
		t.Errorf("error should be DeadlineExceeded, got %v", err)
	}

	cancel()
	if err := ctx.Err(); err != DeadlineExceeded {
		t.Errorf("error should still be DeadlineExceeded, got %v", err)
	}
}

func TestWithValue(t *testing.T) {
	tc := []struct{
		key, val, keyRet, valRet interface{}
		shouldPanic bool
	}{
		{"a", "b", "a", "b", false},
		{"a", "b", "c", nil, false},
		{42, true, 42, true, false},
		{42, true, int64(42), nil, false},
		{nil, true, nil, nil, true},
		{[]int{1, 2, 3}, true, []int{1, 2, 3}, nil, true},
	}

	for _, tt := range tc {
		t.Run("", func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.shouldPanic {
					t.Errorf("should panic")
				}
			}()

			ctx := WithValue(Background(), tt.key, tt.val)

			if tt.valRet != ctx.Value(tt.keyRet) {
				t.Errorf("should receive value %v, got %v", tt.valRet, ctx.Value(tt.keyRet))
			}
		})
	}
}
