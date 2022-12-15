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
	errchanImpl.ch = make(chan struct{})
	errchanImpl.setActive(true)
	return errchanImpl.ch
}

func New(err error) {
	if err == nil || errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return
	}

	errsetImpl.mtx.Lock()
	errsetImpl.errors = append(errsetImpl.errors, err)
	errsetImpl.mtx.Unlock()

	if errchanImpl.isActive() {
		close(errchanImpl.ch)
		errchanImpl.setActive(false)
	}
}

func Error() error {
	errsetImpl.mtx.Lock()
	defer errsetImpl.mtx.Unlock()

	if len(errsetImpl.Errors()) == 0 {
		return nil
	}

	return errsetImpl
}

type errchan struct {
	mtx    sync.Mutex
	active bool
	ch     chan struct{}
}

func (e *errchan) setActive(active bool) {
	e.mtx.Lock()
	defer e.mtx.Unlock()

	e.active = active
}

func (e *errchan) isActive() bool {
	e.mtx.Lock()
	defer e.mtx.Unlock()

	return e.active
}

type errset struct {
	mtx    sync.Mutex
	errors []error
}

func (e *errset) Error() (res string) {
	e.mtx.Lock()
	defer e.mtx.Unlock()

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

func (e *errset) Errors() []error {
	return e.errors
}

func (e *errset) Is(target error) bool {
	e.mtx.Lock()
	defer e.mtx.Unlock()

	for _, err := range e.errors {
		if errors.Is(err, target) {
			return true
		}
	}

	return false
}

func (e *errset) As(target interface{}) bool {
	e.mtx.Lock()
	defer e.mtx.Unlock()

	for _, err := range e.errors {
		if errors.As(err, target) {
			return true
		}
	}

	return false
}
