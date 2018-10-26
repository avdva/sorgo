package async

import (
	"errors"
	"time"
)

var (
	// TimeoutErr indicates, that a timeout has elapsed.
	TimeoutErr = errors.New("timeout")
)

// DoWithTimeout runs f in a different goroutine
//	if f returns before timeout elapses, doAsyncTimeout returns the result of f().
//	otherwise it returns "timeout" error, and calls tmFunc after f returns.
func DoWithTimeout(f func() error, tmFunc func(error), timeout time.Duration) error {
	errs := make(chan error)
	go func() {
		err := f()
		select {
		case errs <- err:
		default:
			if tmFunc != nil {
				tmFunc(err)
			}
		}
	}()
	select {
	case err := <-errs:
		return err
	case <-time.After(timeout):
		return TimeoutErr
	}
}
