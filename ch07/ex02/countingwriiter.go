package main

import (
	"io"
	"log"
)

type wrapWriter struct {
	writer io.Writer
	count  *int64
}

func (c wrapWriter) Write(p []byte) (int, error) {
	log.Printf("before: %v\n", c.count)
	n, err := c.writer.Write(p)
	if err == nil {
		*c.count += int64(n)
	}
	log.Printf("after: %v\n", c.count)
	return n, err
}

// CountingWriter Writerと書き込んだバイト数を持つstructureでwrap
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	wrapper := wrapWriter{w, new(int64)}
	return wrapper, wrapper.count
}
