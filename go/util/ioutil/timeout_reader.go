package ioutil

import (
	"errors"
	"io"
	"time"
)

type readResult struct {
	n     int
	error error
}

type TimeoutReader struct {
	reader  io.ReadCloser
	timeout time.Duration
}

func NewTimeoutReader(reader io.ReadCloser, timeout time.Duration) *TimeoutReader {
	return &TimeoutReader{reader: reader, timeout: timeout}
}

var ErrTimeout = errors.New("timeout")

func (t *TimeoutReader) Read(p []byte) (n int, err error) {
	reads := make(chan readResult, 1)
	go func() {
		n, err := t.reader.Read(p)
		reads <- readResult{n, err}
	}()
	timer := time.NewTimer(t.timeout)
	select {
	case result := <-reads:
		if !timer.Stop() {
			<-timer.C
		}
		return result.n, result.error
	case <-timer.C:
		return 0, ErrTimeout
	}
}

func (t *TimeoutReader) Close() error {
	return t.reader.Close()
}
