package main

import (
	"io"
)

type CountWriter struct {
	n int64
	w io.Writer
}

func (c *CountWriter) Write(data []byte) (int, error) {
	c.n += int64(len(data))
	return c.w.Write(data)
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	newWriter := &CountWriter{
		n: 0,
		w: w,
	}
	return newWriter, &(newWriter.n)
}
