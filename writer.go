// Author seth.a.hoenig@gmail.com
// License MIT

// Package counting provides a writer that counts
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

func NewCountingWriter(w io.Writer) *CountingWriter {
	return &CountingWriter{w, 0}
}

func (w *CountingWriter) Write(p []byte) (n int, e error) {
	n, e = w.writer.Write(p)
	w.written += n
	return n, e
}

func (w *CountingWriter) Written() int {
	return w.written
}
