package whitebackground

import (
	"image"
	_ "image/jpeg"
	"log"
	"net/http"
	"runtime"
	"testing"
)

const url string = "https://d2gg9evh47fn9z.cloudfront.net/800px_COLOURBOX3871999.jpg"

func TestHasWhiteBackground(t *testing.T) {
	runtime.GOMAXPROCS(4)
	// don't worry about errors
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}

	defer response.Body.Close()

	i, _, err := image.Decode(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	if !HasWhiteBackground(i) {
		t.Fatal("Failed")
	}
}

func BenchmarkHasWhiteBackground_1CPU(b *testing.B) {
	benchmarkHasWhiteBackgroundCore(1, b)
}

func BenchmarkHasWhiteBackground_2CPU(b *testing.B) {
	benchmarkHasWhiteBackgroundCore(2, b)
}

func BenchmarkHasWhiteBackground_3CPU(b *testing.B) {
	benchmarkHasWhiteBackgroundCore(3, b)
}

func BenchmarkHasWhiteBackground_4CPU(b *testing.B) {
	benchmarkHasWhiteBackgroundCore(4, b)
}

func benchmarkHasWhiteBackgroundCore(cpus int, b *testing.B) {
	runtime.GOMAXPROCS(cpus)

	// don't worry about errors
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}

	defer response.Body.Close()

	i, _, err := image.Decode(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		HasWhiteBackground(i)
	}
}
