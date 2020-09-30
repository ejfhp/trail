package writer

import (
	"fmt"
	"os"
	"sync"
)

type WriterLimited struct {
	cur     string
	old     string
	curf    *os.File
	limit   int
	written int
}

//NewWriterLimited create a new writer that writes to currentlog copying the old file to oldlog when limitbytes is reach.
func NewWriterLimited(currentlog, oldlog string, limitbytes int) (*WriterLimited, error) {
	w := &WriterLimited{cur: currentlog, old: oldlog, limit: limitbytes, written: -1}
	err := w.Update()
	return w, err
}

func (w *WriterLimited) Write(p []byte) (int, error) {
	mx := sync.Mutex{}
	mx.Lock()
	if err := w.Update(); err != nil {
		mx.Unlock()
		return -1, err
	}
	i, err := w.curf.Write(p)
	if err != nil {
		fmt.Printf("cannot write n:%d %v %v", i, p, err)
	}
	w.written = w.written + i
	mx.Unlock()
	if err != nil {
		return -1, err
	}
	return i, nil
}

func (w *WriterLimited) Update() error {
	switch {
	case w.written < 0:
		return w.reopen()
	case w.written < w.limit:
		return nil
	default:
		w.written = 0
		return w.reopen()
	}
}

func (w *WriterLimited) reopen() error {
	if w.curf != nil {
		if err := w.curf.Close(); err != nil {
			return fmt.Errorf("cannot close current file %s:%w", w.cur, err)
		}
	}
	if err := os.Rename(w.cur, w.old); err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("cannot rename current file %s to %s:%w", w.cur, w.old, err)
		}
	}
	f, err := os.OpenFile(w.cur, os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		return fmt.Errorf("cannot open log file %s:%w", w.cur, err)
	}
	w.curf = f
	return nil
}
