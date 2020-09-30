package writer_test

import (
	"os"
	"testing"
	"time"

	"github.com/ejfhp/trail/writer"
)

func TestNewWriterAMPM(t *testing.T) {
	am := "/tmp/log_am.log"
	pm := "/tmp/log_pm.log"
	text := "most users will use open or create instead.\n"
	w, err := writer.NewWriterAMPM(am, pm)
	if err != nil {
		t.Fatalf("cannot create am/pm writer %v", err)
	}
	var amsizea, amsizeb, pmsizea, pmsizeb int64 = 0, 0, 0, 0
	amia, err := os.Stat(am)
	if err == nil && amia != nil {
		amsizea = amia.Size()
	}
	pmia, err := os.Stat(pm)
	if err == nil && pmia != nil {
		pmsizea = pmia.Size()
	}

	_, err = w.Write([]byte(text))
	if err != nil {
		t.Fatalf("cannot write %v", err)
	}

	amib, err := os.Stat(am)
	if err == nil && amib != nil {
		amsizeb = amib.Size()
	}
	pmib, err := os.Stat(pm)
	if err == nil && pmib != nil {
		pmsizeb = pmib.Size()
	}
	amdiff := amsizeb - amsizea
	pmdiff := pmsizeb - pmsizea
	if amdiff+pmdiff != int64(len(text)) {
		t.Fatalf("text has not been written on the log files %d!=%d", amdiff+pmdiff, len(text))
	}
}

func BenchmarkWriterAMPM(b *testing.B) {
	am := "/tmp/log_am.log"
	pm := "/tmp/log_pm.log"
	text1 := "TEXT 1\n"
	text2 := "TEXT 2\n"
	text3 := "TEXT 3\n"
	text4 := "TEXT 4\n"
	text5 := "TEXT 5\n"
	w, err := writer.NewWriterAMPM(am, pm)
	if err != nil {
		b.Fatalf("cannot create am/pm writer %v", err)
	}
	t := time.Now()
	b.ResetTimer()
	go func() {
		for i := 0; i < b.N; i++ {
			w.UpdateAMPM(t.Add(time.Duration(10 * time.Hour)))
			w.Write([]byte(text1))
		}
	}()
	go func() {
		for i := 0; i < b.N; i++ {
			w.UpdateAMPM(t.Add(time.Duration(4 * time.Hour)))
			w.Write([]byte(text2))
		}
	}()
	go func() {
		for i := 0; i < b.N; i++ {
			w.UpdateAMPM(t.Add(time.Duration(15 * time.Hour)))
			w.Write([]byte(text3))
		}
	}()
	go func() {
		for i := 0; i < b.N; i++ {
			w.Write([]byte(text4))
		}
	}()
	go func() {
		for i := 0; i < b.N; i++ {
			w.UpdateAMPM(t.Add(time.Duration(13 * time.Hour)))
			w.Write([]byte(text5))
		}
	}()
	b.ReportAllocs()
}
