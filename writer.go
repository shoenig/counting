// Author seth.a.hoenig@gmail.com
// License MIT

// Package counting provides an io.Writer that passively counts bytes.
package counting

import (
	"io"
)

// CountingWriter passivly counts the number of bytes that pass through it.
// It is not thread safe.
type CountingWriter struct {
	writer  io.Writer
	written int
}

// NewCountingWriter wraps a Writer and keeps track of the number of
// bytes that have been written to the Writer.
func NewCountingWriter(w io.Writer) *CountingWriter {
	return &CountingWriter{w, 0}
}

// Write writes bytes and returns the number of bytes written.
func (w *CountingWriter) Write(p []byte) (n int, e error) {
	n, e = w.writer.Write(p)
	w.written += n
	return n, e
}

// Written returns the number of bytes that have been written so far.
func (w *CountingWriter) Written() int {
	return w.written
}
