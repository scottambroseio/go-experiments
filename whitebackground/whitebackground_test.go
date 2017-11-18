package whitebackground

import (
	"image"
	_ "image/jpeg"
	"log"
	"net/http"
	"runtime"
	"testing"
)

// 800 x 800
const url string = "https://d2gg9evh47fn9z.cloudfront.net/800px_COLOURBOX3871999.jpg"

// 5627 × 3517
const bigURL string = "https://www.walldevil.com/wallpapers/w04/137357-artwork-flowers-scars-style-white-background-women-women-artwork.jpg"

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

func BenchmarkHasWhiteBackground_1CPU_800x600(b *testing.B) {
	benchmarkHasWhiteBackgroundCore(1, b, url)
}

func BenchmarkHasWhiteBackground_2CPU_800x600(b *testing.B) {
	benchmarkHasWhiteBackgroundCore(2, b, url)
}

func BenchmarkHasWhiteBackground_3CPU_800x600(b *testing.B) {
	benchmarkHasWhiteBackgroundCore(3, b, url)
}

func BenchmarkHasWhiteBackground_4CPU_800x600(b *testing.B) {
	benchmarkHasWhiteBackgroundCore(4, b, url)
}

func BenchmarkHasWhiteBackground_1CPU_5627x3517(b *testing.B) {
	benchmarkHasWhiteBackgroundCore(1, b, bigURL)
}

func BenchmarkHasWhiteBackground_2CPU_5627x3517(b *testing.B) {
	benchmarkHasWhiteBackgroundCore(2, b, bigURL)
}

func BenchmarkHasWhiteBackground_3CPU_5627x3517(b *testing.B) {
	benchmarkHasWhiteBackgroundCore(3, b, bigURL)
}

func BenchmarkHasWhiteBackground_4CPU_5627x3517(b *testing.B) {
	benchmarkHasWhiteBackgroundCore(4, b, bigURL)
}

func benchmarkHasWhiteBackgroundCore(cpus int, b *testing.B, url string) {
	runtime.GOMAXPROCS(cpus)

	r, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	i, _, err := image.Decode(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		HasWhiteBackground(i)
	}
}
