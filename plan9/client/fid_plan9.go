package client

import "os"

type Fid struct {
	*os.File
}

func (f *Fid) Close() error {
	if f == nil {
		// The os package returns os.ErrInvalid.
		// Return nil for p9p compatibility.
		return nil
	}
	return f.File.Close()
}
