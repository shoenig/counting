// Author seth.a.hoenig@gmail.com

package counting

import (
	"compress/gzip"
	"io/ioutil"
	"os"
	"testing"
)

func Test_WriteCompressed(t *testing.T) {
	f, e := ioutil.TempFile("", "counting-")

	if e != nil {
		t.Fatal("create tempfile failed", e)
	}

	counter := NewCountingWriter(f)

	compressor, e := gzip.NewWriterLevel(counter, 9)
	if e != nil {
		t.Fatal("failed to create compressor", e)
	}

	// there are 42 a's
	s := []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	n, e := compressor.Write(s)

	if e != nil {
		t.Fatal("compressed write failed", e)
	}

	if n != len(s) {
		t.Fatal("compressed file is missing bytes")
	}

	e = compressor.Close()
	if e != nil {
		t.Fatal("compressor failed to close", e)
	}

	bytes := counter.Written()
	if bytes != 27 {
		t.Fatal("counter counted incorrectly")
	}

	// cleanup
	e = os.RemoveAll(f.Name())
	if e != nil {
		t.Fatal("failed to cleanup", e)
	}
}
