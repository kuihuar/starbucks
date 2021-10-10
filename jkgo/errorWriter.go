package jkgo

import (
	"fmt"
	"io"
)

type errWriter struct {
	io.Writer
	err error
}

type Status struct {
	Code int
	Reason string
}
type Header struct {
	Key string
	Value string
}
//消除Errors
func (e *errWriter)Write(buf []byte) (int, error) {
	if e.err != nil {
		return 0, e.err
	}
	var n int
	n, e.err = e.Writer.Write(buf)
	return n, nil
}
//
func WriteResponse(w io.Writer, st Status, h []Header, body io.Reader)error  {
	ew := &errWriter{Writer:w}
	fmt.Fprintf(ew, "HTTP/1.1 %d %s\t\n", st.Code, st.Reason)
	for _,h := range h {
		fmt.Fprintf(ew, "%s: %s\r\n", h.Key, h.Value)
	}
	fmt.Fprintf(ew, "\r\n")
	return ew.err
}