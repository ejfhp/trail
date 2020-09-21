package trace

import (
	"time"
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
	timef          = "time"
)

//Trace is a log record. It has to be built with a chain of calls and finalized calling Finalize()
type Trace []byte

//New returns a new Trace record
func New() Trace {
	return append(make([]byte, 0, 500), '{')
}

func withSeverity(s, msg string) Trace {
	r := New()
	r.appendBytes(severity, s)
	r.appendBytes(message, msg)
	return r
}

//Debug start a corresponding severity Trace record
func Debug(msg string) Trace {
	return withSeverity(debug, msg)
}

//Info start a corresponding severity Trace record
func Info(msg string) Trace {
	return withSeverity(info, msg)
}

//Warning start a corresponding severity Trace record
func Warning(msg string) Trace {
	return withSeverity(warning, msg)
}

//Alert start a corresponding severity Trace record
func Alert(msg string) Trace {
	return withSeverity(alert, msg)
}

//Add append a generic key-value to the current Trace
func (r Trace) Add(name, val string) Trace {
	r.appendBytes(name, val)
	return r
}

//Source append a source location to the current Trace
func (r Trace) Source(file, function, lineNum string) Trace {
	r.appendBytesNew(source)
	r.appendBytes(sourceFile, file)
	r.appendBytes(sourceFunction, function)
	r.appendBytes(sourceLine, lineNum)
	r.close()
	r.appendByte(',')
	return r
}

//Error append an error message to the current Trace
func (r Trace) Error(e error) Trace {
	r.appendBytes(err, e.Error())
	return r
}

//Append another Trace record to the current one
func (r Trace) Append(trail Trace) Trace {
	return append(r, trail[1:]...)
}

//UTC add the time to the current Trace record
func (r Trace) UTC() Trace {
	r.appendBytes(timef, time.Now().Format(time.RFC3339))
	return r
}

//Finalize returns the built Trace record as JSON
func (r Trace) Finalize() []byte {
	r.close()
	return r
}

func (r *Trace) appendBytes(key, val string) {
	*r = append(*r, '"')
	*r = append(*r, []byte(key)...)
	*r = append(*r, '"', ':', '"')
	*r = append(*r, []byte(val)...)
	*r = append(*r, '"', ',')
}

func (r *Trace) appendBytesNew(key string) {
	*r = append(*r, '"')
	*r = append(*r, []byte(key)...)
	*r = append(*r, '"', ':', '{')
}

func (r *Trace) appendByte(b ...byte) {
	*r = append(*r, b...)
}

func (r *Trace) close() {
	(*r)[len(*r)-1] = '}'
}
