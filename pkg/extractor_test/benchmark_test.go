package benchmark_test

import (
	"rare/pkg/extractor"
	"testing"
)

func batchInputGenerator(batches int, batchSize int) <-chan []extractor.BString {
	c := make(chan []extractor.BString, 128)
	go func() {
		for i := 0; i < batches; i++ {
			batch := make([]extractor.BString, batchSize)
			for j := 0; j < batchSize; j++ {
				batch[j] = extractor.BString("abcdefg 123")
			}
			c <- batch
		}
		close(c)
	}()
	return c
}

func BenchmarkExtractor(b *testing.B) {
	total := 0
	for n := 0; n < b.N; n++ {
		gen := batchInputGenerator(10000, 100)
		extractor, _ := extractor.New(gen, &extractor.Config{
			Regex:   `(\d{3})`,
			Extract: "{bucket {1} 10}",
			Workers: 2,
		})
		reader := extractor.ReadChan()
		for val := range reader {
			total++
			if val[0].Extracted != "120" {
				panic("NO MATCH")
			}
		} // Drain reader
	}
}
