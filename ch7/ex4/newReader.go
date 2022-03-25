package main

type StringReader string

func NewReader() *StringReader {
	return new(StringReader)
}

func (r *StringReader) Read(b []byte) (int, error) {
	*r += StringReader(b)
	return len(b), nil
}

func (r *StringReader) String() string {
	return string(*r)
}
