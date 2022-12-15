package errset

import (
	"context"
	"errors"
	"sync"
)

var errchanImpl *errchan
var errsetImpl *errset

func init() {
	errchanImpl = &errchan{}
	errsetImpl = &errset{}
}

func Catch() <-chan struct{} {
	return errchanImpl.newChan()
}

func New(err error) {
	if err == nil || errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return
	}

	errsetImpl.add(err)
	errchanImpl.notify()
}

func Error() error {
	errsetImpl.mtx.RLock()
	defer errsetImpl.mtx.RUnlock()

	if len(Errors()) == 0 {
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
	mtx    sync.RWMutex
	active bool
	ch     chan struct{}
	chans  []chan struct{}
}

func (e *errchan) newChan() <-chan struct{} {
	e.mtx.Lock()
	defer e.mtx.Unlock()

	ch := make(chan struct{})
	e.chans = append(e.chans, ch)

	if !e.active {
		e.ch = make(chan struct{})
		e.active = true

		go e.broadcaster()
	}

	return ch
}

func (e *errchan) notify() {
	e.mtx.RLock()
	defer e.mtx.RUnlock()

	if e.active {
		e.ch <- struct{}{}
	}
}

func (e *errchan) broadcaster() {
	<-e.ch

	e.mtx.Lock()
	defer e.mtx.Unlock()

	for _, v := range e.chans {
		close(v)
	}

	e.chans = nil
	e.active = false
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
