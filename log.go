package trail

import (
	"io"
	"sync"
)

const (
	debug          = "DEBUG"
	info           = "INFO"
	warning        = "WARNING"
	alert          = "ALERT"
	severity       = "severity"
	source         = "sourceLocation"
	sourceFile     = "file"
	sourceFunction = "function"
	sourceLine     = "line"
	message        = "message"
	err            = "error"
)

var li *logger
var mu sync.Mutex

type logger struct {
	writer io.Writer
}

func Println(tr []byte) {
	if li != nil {
		writeN(li.writer, tr)
	}
}

func SetMultiwriters(outputs []io.Writer) {
	ws := io.MultiWriter(outputs...)
	l := logger{
		writer: ws,
	}
	li = &l
}

func SetWriter(w io.Writer) {
	l := logger{
		writer: w,
	}
	li = &l
}

func writeN(out io.Writer, data []byte) error {
	mu.Lock()
	defer mu.Unlock()
	_, err := out.Write(data)
	_, err = out.Write([]byte{'\n'})
	return err
}
