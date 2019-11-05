package splitsies_test

import (
	"fmt"
	"testing"

	splitsies "github.com/saadullahsaeed/splitsies/lib"
)

func BenchmarkSplitter_Split(b *testing.B) {
	splitter := splitsies.New("test_", 1)
	for n := 0; n < b.N; n++ {
		_, err := splitter.Split("../testdata/transaction.csv", "test_")
		if err != nil {
			b.Error(err)
		}
	}
}

func TestSplitter_SplitWithoutReuse(t *testing.T) {
	splitter := splitsies.New("test_", 1)
	_, err := splitter.Split("../testdata/test.csv", "test_")
	fmt.Println(err)
}
