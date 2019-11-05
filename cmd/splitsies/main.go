package main

import (
	"flag"
	"log"
	"os"

	splitsies "github.com/saadullahsaeed/splitsies/lib"
)

func main() {
	hasHeader := flag.Bool("header", true, "Indicates if the input file has a header")
	fnPrefix := flag.String("out-file-prefix", "", "The prefix to use for filenames generated for output")
	fnColIndex := flag.Int("out-file-col-index", 0, "The column index from the input file to use in the generated file name (e.g. date column) - uses first column by default")
	cvMaxLength := flag.Int("out-file-col-max-length", 0, "Maximum length of the value in the specified column (by default no limit i.e. 0)")
	inCSV := flag.String("file", "", "Input file to split")
	outputDir := flag.String("out-dir", "", "Directory to write the generated files to")

	flag.Parse()

	if os.Args[1] == "help" || os.Args[1] == "usage" {
		flag.Usage()
		return
	}

	splitter := splitsies.Splitter{
		WithHeader:           *hasHeader,
		FileNamePrefix:       *fnPrefix,
		FileNameColumnIndex:  *fnColIndex,
		ColumnValueMaxLength: *cvMaxLength,
		InputCSVPath:         *inCSV,
		OutputDirPath:        *outputDir,
	}
	_, err := splitter.Split()
	if err != nil {
		log.Fatal(err)
	}
}
