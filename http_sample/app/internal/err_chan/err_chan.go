package err_chan

import (
	"context"
	"errors"
	"sync"
)

type err struct {
	mtx sync.Mutex
	err error
	ch  chan struct{}
}

var errImpl err

func init() {
	errImpl = err{
		ch: make(chan struct{}),
	}
}

func Catch() <-chan struct{} {
	return errImpl.ch
}

func New(err error) {
	if err == nil {
		return
	}

	errImpl.mtx.Lock()
	defer errImpl.mtx.Unlock()
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return
	}

	errImpl.err = err
	errImpl.ch <- struct{}{}
}

func Error() error {
	return errImpl.err
}
