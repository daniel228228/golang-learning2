package errset

import (
	"context"
	"errors"
	"sync"
)

var errchanImpl *errchan
var errsetImpl *errset

func init() {
	errchanImpl = &errchan{
		ch: make(chan error),
	}
	errsetImpl = &errset{}
}

func Catch() <-chan error {
	return errchanImpl.newChan()
}

func New(err error) {
	if err == nil || errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return
	}

	errsetImpl.add(err)
	errchanImpl.notify(err)
}

func Error() error {
	errsetImpl.mtx.RLock()
	defer errsetImpl.mtx.RUnlock()

	if len(errsetImpl.errors) == 0 {
		return nil
	}

	return errsetImpl
}

func Errors() []error {
	errsetImpl.mtx.RLock()
	defer errsetImpl.mtx.RUnlock()

	return errsetImpl.errors
}

type errchan struct {
	mtx   sync.Mutex
	ch    chan error
	chans []chan error
}

func (e *errchan) newChan() <-chan error {
	ch := make(chan error, 1)

	e.mtx.Lock()
	defer e.mtx.Unlock()

	e.chans = append(e.chans, ch)

	if len(e.chans) == 1 {
		go e.broadcaster()
	}

	return ch
}

func (e *errchan) notify(err error) {
	e.mtx.Lock()
	active := len(e.chans) > 0
	e.mtx.Unlock()

	if active {
		e.ch <- err
	}
}

func (e *errchan) broadcaster() {
	err := <-e.ch

	e.mtx.Lock()
	defer e.mtx.Unlock()

	for _, v := range e.chans {
		v <- err
		close(v)
	}

	e.chans = nil
}

type errset struct {
	mtx    sync.RWMutex
	errors []error
}

func (e *errset) add(err error) {
	e.mtx.Lock()
	defer e.mtx.Unlock()

	e.errors = append(e.errors, err)
}

func (e *errset) Error() (res string) {
	e.mtx.RLock()
	defer e.mtx.RUnlock()

	switch len(e.errors) {
	case 0:
		return ""
	case 1:
		return e.errors[0].Error()
	}

	for i := 0; i < len(e.errors)-1; i++ {
		res += e.errors[i].Error() + "; "
	}

	res += e.errors[len(e.errors)-1].Error()

	return
}

func (e *errset) Is(target error) bool {
	e.mtx.RLock()
	defer e.mtx.RUnlock()

	for _, err := range e.errors {
		if errors.Is(err, target) {
			return true
		}
	}

	return false
}

func (e *errset) As(target interface{}) bool {
	e.mtx.RLock()
	defer e.mtx.RUnlock()

	for _, err := range e.errors {
		if errors.As(err, target) {
			return true
		}
	}

	return false
}
