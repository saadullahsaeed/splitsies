package splitsies_test

import (
	"io/ioutil"
	"os"
	"testing"

	splitsies "github.com/saadullahsaeed/splitsies/lib"
)

var generatedFiles = map[string]string{
	"test_1.csv": "1,a,a",
	"test_2.csv": "2,b,a",
	"test_3.csv": "3,c,a",
	"test_4.csv": "4,d,a",
	"test_5.csv": "5,e,a",
}

func setup() {
	for file := range generatedFiles {
		file = "../testdata/" + file
		_, err := os.Stat(file)
		if os.IsNotExist(err) {
			continue
		}
		os.Remove(file)
	}
}

func BenchmarkSplitter_Split(b *testing.B) {
	setup()

	splitter := splitsies.Splitter{
		WithHeader:          true,
		FileNamePrefix:      "test_",
		FileNameColumnIndex: 0,
		InputCSVPath:        "../testdata/test.csv",
		OutputDirPath:       "../testdata/",
	}
	for n := 0; n < b.N; n++ {
		_, err := splitter.Split()
		if err != nil {
			b.Error(err)
		}
	}
}

func TestSplitter_Split(t *testing.T) {
	defer setup()

	splitter := splitsies.Splitter{
		WithHeader:          true,
		FileNamePrefix:      "test_",
		FileNameColumnIndex: 0,
		InputCSVPath:        "../testdata/test.csv",
		OutputDirPath:       "../testdata/",
	}
	_, err := splitter.Split()
	if err != nil {
		t.Error(err)
		return
	}

	for file, value := range generatedFiles {
		file = "../testdata/" + file
		_, err := os.Stat(file)
		if os.IsNotExist(err) {
			t.FailNow()
		}

		content, _ := ioutil.ReadFile(file)
		if string(content) != value+"\n" {
			t.Errorf("expected %s but got %s", value, string(content))
		}
	}
}
