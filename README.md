## Counting

Provides CountingWriter which counts the number of bytes passing through it

### Example

    package main

    import (
    	"compress/gzip"
	    "fmt"
	    "io/ioutil"

	    "github.com/shoenig/counting"
    )

    func main() {
	    f, _ := ioutil.TempFile("", "counting-")
	    counter := counting.NewCountingWriter(f)
	    compressor, _ := gzip.NewWriterLevel(counter, 9)
	    s := []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa") // 42 bytes
	    u, _ := compressor.Write(s)
	    compressor.Close()
	    c := counter.Written()
	    fmt.Println("uncompressed bytes written", u)
	    fmt.Println("compressed bytes written", c)
    }


