package writer

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type WriterAMPM struct {
	amf, pmf string
	curf     *os.File
	am       bool
}

//NewWriterAMPM create a new writer that write in am during the morning and to
//pm during the afternoon.
func NewWriterAMPM(am, pm string) (*WriterAMPM, error) {
	wam := &WriterAMPM{amf: am, pmf: pm}
	wam.UpdateAMPM(time.Now())
	err := wam.reopen()
	return wam, err
}

func (w *WriterAMPM) Write(p []byte) (int, error) {
	toUpdate := w.UpdateAMPM(time.Now())
	mx := sync.Mutex{}
	mx.Lock()
	if toUpdate {
		if err := w.reopen(); err != nil {
			mx.Unlock()
			return -1, err
		}
	}
	i, err := w.curf.Write(p)
	mx.Unlock()
	if err != nil {
		return -1, err
	}
	return i, nil
}

func (w *WriterAMPM) UpdateAMPM(t time.Time) bool {
	t24, err := strconv.Atoi(t.Format("15"))
	if err != nil {
		t24 = 7
	}
	am := false
	if t24 < 12 {
		am = true
	}
	if w.am != am {
		w.am = am
		return true
	}
	return false
}

func (w *WriterAMPM) reopen() error {
	if w.curf != nil {
		if err := w.curf.Close(); err != nil {
			return fmt.Errorf("cannot close current file %s:%w", w.amf, err)
		}
	}
	switch w.am {
	case true:
		f, err := os.OpenFile(w.amf, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0664)
		if err != nil {
			return fmt.Errorf("cannot open am file %s:%w", w.pmf, err)
		}
		w.curf = f
	case false:
		f, err := os.OpenFile(w.pmf, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0664)
		if err != nil {
			return fmt.Errorf("cannot open am file %s:%w", w.pmf, err)
		}
		w.curf = f
	}
	return nil
}
