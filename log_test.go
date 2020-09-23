package trail_test

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/ejfhp/trail"
	"github.com/ejfhp/trail/trace"
)

func TestLog(t *testing.T) {
	output := strings.Builder{}
	msg := trace.Info("failed to log")
	msga := trace.New().Add("extra", "append")
	trail.SetWriter(&output)

	trail.Println(msg.Append(msga))

	exp := `{"severity":"INFO","message":"failed to log","extra":"append"}` + "\n"
	log := output.String()
	if log != exp {
		t.Fatalf("unexpected output:\n'%s' instead of:\n'%s'", log, exp)
	}
}

func TestSourceLog(t *testing.T) {
	output := strings.Builder{}
	msg := trace.Info("failed to log").Source("log_test.go", "TestSourceLog", "99")
	trail.SetWriter(&output)

	trail.Println(msg)

	exp := `{"severity":"INFO","message":"failed to log","sourceLocation":{"file":"log_test.go","function":"TestSourceLog","line":"99"}}` + "\n"
	log := output.String()
	if log != exp {
		t.Fatalf("unexpected output:\n'%s' instead of:\n'%s'", log, exp)
	}
	fmt.Println(string(msg))
}

func BenchmarkLog(b *testing.B) {
	trail.SetWriter(ioutil.Discard)
	// err := fmt.Errorf("E")
	num := "9"
	function := "B"
	file := "l"
	r := trace.Info("ciao")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// trail.Println(trail.New().Append(r).Source(file, function, num).Error(err).Add("x1", "y1").Add("x2", "y2").Add("x3", "y3").Add("x4", "y4").Add("x5", "y5"))
		trail.Println(trace.New().Append(r).Source(file, function, num).Add("x1", "y1").Add("x2", "y2").Add("x3", "y3").Add("x4", "y4").Add("x5", "y5"))
	}
	b.ReportAllocs()
}

func BenchmarkTimeLog(b *testing.B) {
	trail.SetWriter(ioutil.Discard)
	// err := fmt.Errorf("E")
	num := "9"
	function := "B"
	file := "l"
	r := trace.Info("ciao")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// trail.Println(trail.New().Append(r).Source(file, function, num).Error(err).Add("x1", "y1").Add("x2", "y2").Add("x3", "y3").Add("x4", "y4").Add("x5", "y5"))
		trail.Println(trace.New().UTC().Append(r).Source(file, function, num).Add("x1", "y1").Add("x2", "y2").Add("x3", "y3").Add("x4", "y4").Add("x5", "y5"))
	}
	b.ReportAllocs()
}
