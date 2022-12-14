package errset

import (
	"context"
	"errors"
	"sync"
)

var Ch chan struct{}
var errsetImpl *errset

func init() {
	Ch = make(chan struct{})
	errsetImpl = &errset{}
}

func New(err error) {
	if err == nil || errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return
	}

	errsetImpl.mtx.Lock()
	defer errsetImpl.mtx.Unlock()
	errsetImpl.errors = append(errsetImpl.errors, err)
	Ch <- struct{}{}
}

func Error() error {
	errsetImpl.mtx.Lock()
	defer errsetImpl.mtx.Unlock()

	if len(errsetImpl.Errors()) == 0 {
		return nil
	}

	return errsetImpl
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
