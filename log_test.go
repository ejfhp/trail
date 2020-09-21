package trail_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/ejfhp/gstl/trail"
	"github.com/ejfhp/gstl/trail/trace"
)

func TestLog(t *testing.T) {
	output := strings.Builder{}
	msg := trace.Info("failed to log")
	msga := trace.New().Add("extra", "append")
	trail.SetWriter(&output)

	trail.Println(msg.Append(msga).Finalize())

	exp := `{"severity":"INFO","message":"failed to log","extra":"append"}` + "\n"
	log := output.String()
	if log != exp {
		t.Fatalf("unexpected output:\n'%s' instead of:\n'%s'", log, exp)
	}
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
