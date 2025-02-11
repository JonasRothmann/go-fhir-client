package fhirclient

import (
	"bytes"
	"io"
)

type teeReadCloser struct {
	io.Reader
	closer io.Closer
	buf    *bytes.Buffer
}

func newTeeReadCloser(rc io.ReadCloser, buf *bytes.Buffer) io.ReadCloser {
	return &teeReadCloser{
		Reader: io.TeeReader(rc, buf),
		closer: rc,
		buf:    buf,
	}
}

func (t *teeReadCloser) Close() error {
	err := t.closer.Close()

	return err
}
