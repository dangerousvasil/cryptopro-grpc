package rwCloser

import "io"

const DEFAULT_BUFFER_LENGTH = 4096

type RWCloser struct {
	io.ReadCloser
	io.WriteCloser
}

func (rw RWCloser) Close() error {
	err := rw.ReadCloser.Close()
	if err != nil {
		return err
	}
	if err := rw.WriteCloser.Close(); err != nil {
		return err
	}
	return nil
}
