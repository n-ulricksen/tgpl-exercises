package main

import "io"

type limitReader struct {
	r         io.Reader
	limit     int64
	bytesRead int64
}

// LimitReader wraps an io.Reader with a new io.Reader that reports EOF after
// n bytes.
func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{
		r:     r,
		limit: n,
	}
}

func (lr *limitReader) Read(b []byte) (int, error) {
	// Read bytes as long as the limit has not been reached
	n := int64(len(b))

	if (lr.bytesRead + n) >= lr.limit {
		// Limit will be reached after reading partial slice
		toRead := lr.limit - lr.bytesRead
		lr.r.Read(b[:toRead])
		lr.bytesRead += toRead

		return int(toRead), io.EOF
	}

	// Limit not yet reached, read the whole slice
	lr.r.Read(b)
	lr.bytesRead += n
	return int(n), nil
}
