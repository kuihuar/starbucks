package interviewquestion

import (
	"errors"
	"time"
)

func DoWithTimeout(f func() error, d time.Duration) error {
	done := make(chan error, 1)
	go func() {
		err := f()
		done <- err
	}()
	select {
	case err := <-done:
		return err
	case <-time.After(d):
		return errors.New("timeout")
	}
}
func DoWithTimeout1(f func() error, d time.Duration) error {
	ch := make(chan error, 1)
	go time.AfterFunc(d, func() {
		ch <- errors.New("timeout")
	})
	go func() {
		ch <- f()
	}()
	return <-ch
}