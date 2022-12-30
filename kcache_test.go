package kv

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

type tableTest struct {
	input string
}

func TestIn(t *testing.T) {
	cache, err := LoadCache(strings.NewReader(`
aaaaaaaaaaa
bbbbbbbbbbb
ccccccccccc`[1:]))
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	if !cache.In("aaaaaaaaaaa") {
		t.Errorf("'aaaaaaaaaaa' not in cache")
	}
	if cache.In("ddddddddddd") {
		t.Errorf("'ddddddddddd' in cache")
	}
}

func BenchmarkIn(b *testing.B) {
	cache, err := LoadCacheFromFile("video_ids_sorted.txt")
	if err != nil {
		b.Fatalf("%v\n", err)
	}

	f, _ := os.Open("random_sample.txt")
	scanner := bufio.NewScanner(f)
	var tt []tableTest
	for i := 0; i < 10 && scanner.Scan(); i++ {
		tt = append(tt, tableTest{scanner.Text()})
	}

	for _, t := range tt {
		b.Run(fmt.Sprintf("input_%s", t.input), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				cache.In(t.input)
			}
		})
	}
}
