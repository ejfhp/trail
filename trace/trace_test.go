package trace_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/ejfhp/trail/trace"
)

func TestTrace(t *testing.T) {
	msg := trace.Info("failed to log")

	exp := `{"severity":"INFO","message":"failed to log"}`
	log := string(msg)
	if log != exp {
		t.Fatalf("unexpected output:\n'%s' instead of:\n'%s'", log, exp)
	}
}

func TestAppendTrace(t *testing.T) {
	msg := trace.Info("failed to log")
	msga := trace.New().Add("extra", "append")
	log := string(msg.Append(msga))

	exp := `{"severity":"INFO","message":"failed to log","extra":"append"}`
	if log != exp {
		t.Fatalf("unexpected output:\n'%s' instead of:\n'%s'", log, exp)
	}
}
func TestAppendInt64Trace(t *testing.T) {
	msg := trace.Info("failed to log")
	msga := trace.New().AddInt64("extra", 20)
	log := string(msg.Append(msga))

	exp := `{"severity":"INFO","message":"failed to log","extra":"20"}`
	if log != exp {
		t.Fatalf("unexpected output:\n'%s' instead of:\n'%s'", log, exp)
	}
}

func TestAppendInt32Trace(t *testing.T) {
	msg := trace.Info("failed to log")
	msga := trace.New().AddInt32("extra", 20)
	log := string(msg.Append(msga))

	exp := `{"severity":"INFO","message":"failed to log","extra":"20"}`
	if log != exp {
		t.Fatalf("unexpected output:\n'%s' instead of:\n'%s'", log, exp)
	}
}

func TestAppendIntTrace(t *testing.T) {
	msg := trace.Info("failed to log")
	msga := trace.New().AddInt("extra", 20)
	log := string(msg.Append(msga))

	exp := `{"severity":"INFO","message":"failed to log","extra":"20"}`
	if log != exp {
		t.Fatalf("unexpected output:\n'%s' instead of:\n'%s'", log, exp)
	}
}
func TestAppendFloat64Trace(t *testing.T) {
	msg := trace.Info("failed to log")
	msga := trace.New().AddFloat64("extra", 20.1)
	log := string(msg.Append(msga))

	exp := `{"severity":"INFO","message":"failed to log","extra":"20.100000"}`
	if log != exp {
		t.Fatalf("unexpected output:\n'%s' instead of:\n'%s'", log, exp)
	}
}

func TestAppendFloat32Trace(t *testing.T) {
	msg := trace.Info("failed to log")
	msga := trace.New().AddFloat32("extra", 20.1)
	log := string(msg.Append(msga))

	exp := `{"severity":"INFO","message":"failed to log","extra":"20.100000"}`
	if log != exp {
		t.Fatalf("unexpected output:\n'%s' instead of:\n'%s'", log, exp)
	}
}

func TestAppendTimeTrace(t *testing.T) {
	msg := trace.Info("failed to log")
	time := time.Date(2025, time.January, 20, 6, 30, 10, 999, time.UTC)
	msga := trace.New().AddTime("extra", time)
	log := string(msg.Append(msga))

	exp := `{"severity":"INFO","message":"failed to log","extra":"2025-01-20T06:30:10Z"}`
	if log != exp {
		t.Fatalf("unexpected output:\n'%s' instead of:\n'%s'", log, exp)
	}
}

func TestErrorTrace(t *testing.T) {
	_, err := os.Open("non_esiste.txt")
	errL := fmt.Errorf("errore di prova: %w", err)
	msg := trace.Alert("error").Error(errL)

	exp := `{"severity":"ALERT","message":"error","error":"errore di prova: open non_esiste.txt: no such file or directory"}`
	log := string(msg)
	if log != exp {
		t.Fatalf("unexpected output:\n'%s' instead of:\n'%s'", log, exp)
	}
}

func TestSourceTrace(t *testing.T) {
	msg := trace.Info("failed to log").Source("log_test.go", "TestSourceTrace", "99")

	exp := `{"severity":"INFO","message":"failed to log","sourceLocation":{"file":"log_test.go","function":"TestSourceTrace","line":"99"}}`
	log := string(msg)
	if log != exp {
		t.Fatalf("unexpected output:\n'%s' instead of:\n'%s'", log, exp)
	}
	fmt.Println(string(msg))
}

func TestSourceErrorTrace(t *testing.T) {
	err := fmt.Errorf("FAKE ERROR")
	msg := trace.Info("failed to log").Source("log_test.go", "TestSourceTrace", "99").Error(err)

	exp := `{"severity":"INFO","message":"failed to log","sourceLocation":{"file":"log_test.go","function":"TestSourceTrace","line":"99"},"error":"FAKE ERROR"}`
	log := string(msg)
	if log != exp {
		t.Fatalf("unexpected output:\n'%s' instead of:\n'%s'", log, exp)
	}
	fmt.Println(string(msg))
}

func BenchmarkLog(b *testing.B) {
	num := "9"
	function := "B"
	file := "l"
	r := trace.Info("ciao")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// trail.Println(trail.New().Append(r).Source(file, function, num).Error(err).Add("x1", "y1").Add("x2", "y2").Add("x3", "y3").Add("x4", "y4").Add("x5", "y5"))
		trace.New().Append(r).Source(file, function, num).Add("x1", "y1").Add("x2", "y2").Add("x3", "y3").Add("x4", "y4").Add("x5", "y5")
	}
	b.ReportAllocs()
}
