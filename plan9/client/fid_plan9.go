package client

import "os"

type Fid struct {
	*os.File
}

// Close closes the fid.
//
// This wrapper prevents runtime panic if *Fid is nil.
// See https://github.com/golang/go/issues/18617
func (f *Fid) Close() error {
	if f == nil {
		return os.ErrInvalid
	}
	return f.File.Close()
}
