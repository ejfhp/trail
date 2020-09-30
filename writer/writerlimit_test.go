package writer_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/ejfhp/trail/writer"
)

func TestNewWriterLimited(t *testing.T) {
	f1 := "/tmp/log_cur.log"
	f2 := "/tmp/log_old.log"
	text := "most users will use open or create instead.\n"
	limit := 30000
	repeat := 1000
	w, err := writer.NewWriterLimited(f1, f2, limit)
	if err != nil {
		t.Fatalf("cannot create limited writer %v", err)
	}
	var f1sizea, f1sizeb, f2sizea, f2sizeb int64 = 0, 0, 0, 0
	f1a, err := os.Stat(f1)
	if err == nil && f1a != nil {
		f1sizea = f1a.Size()
	}

	f2a, err := os.Stat(f2)
	if err == nil && f2a != nil {
		f2sizea = f2a.Size()
	}
	fmt.Printf("Before f1:%d f2:%d\n", f1sizea, f2sizea)

	for i := 0; i < repeat; i++ {
		_, err = w.Write([]byte(text))
		if err != nil {
			t.Fatalf("cannot write %v", err)
		}
	}

	f1b, err := os.Stat(f1)
	if err == nil && f1b != nil {
		f1sizeb = f1b.Size()
	}

	f2b, err := os.Stat(f2)
	if err == nil && f2b != nil {
		f2sizeb = f2b.Size()
	}
	fmt.Printf("After cur:%d old:%d\n", f1sizeb, f2sizeb)
	if f1sizeb+f2sizeb != int64(len(text)*repeat) {
		t.Fatalf("text has not been written on the log files %d!=%d", f1sizeb+f2sizeb, len(text)*repeat)
	}
}

func BenchmarkWriterLimited(b *testing.B) {
	f1 := "/tmp/log_cur.log"
	f2 := "/tmp/log_old.log"
	limit := 1000
	text1 := "TEXT 1\n"
	text2 := "TEXT 2\n"
	text3 := "TEXT 3\n"
	text4 := "TEXT 4\n"
	text5 := "TEXT 5\n"
	w, err := writer.NewWriterLimited(f1, f2, limit)
	if err != nil {
		b.Fatalf("cannot create limited writer %v\n", err)
	}
	b.ResetTimer()
	go func() {
		for i := 0; i < b.N; i++ {
			_, err := w.Write([]byte(text1))
			if err != nil {
				b.Fatalf("1 cannot write %v\n", err)
			}
		}
	}()
	go func() {
		for i := 0; i < b.N; i++ {
			_, err := w.Write([]byte(text2))
			if err != nil {
				b.Fatalf("2 cannot write %v\n", err)
			}
		}
	}()
	go func() {
		for i := 0; i < b.N; i++ {
			_, err := w.Write([]byte(text3))
			if err != nil {
				b.Fatalf("3 cannot write %v\n", err)
			}
		}
	}()
	go func() {
		for i := 0; i < b.N; i++ {
			_, err := w.Write([]byte(text4))
			if err != nil {
				b.Fatalf("4 cannot write %v\n", err)
			}
		}
	}()
	go func() {
		for i := 0; i < b.N; i++ {
			_, err := w.Write([]byte(text5))
			if err != nil {
				b.Fatalf("5 cannot write %v\n", err)
			}
		}
	}()
	b.ReportAllocs()
}
