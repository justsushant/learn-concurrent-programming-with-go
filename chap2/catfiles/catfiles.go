package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// go run catfiles.go ./testdata/file1.txt ./testdata/file2.txt ./testdata/file3.txt

func main() {
	CatFiles(os.Args[1:], os.Stdout)
}

func CatFiles(fileNames []string, out io.Writer) {
	for _, fileName := range fileNames {
		go readAndPrint(fileName, out)
	}

	time.Sleep(3 * time.Second)
}

func readAndPrint(fileName string, out io.Writer) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintln(out, err.Error())
		return
	}

	b, err := io.ReadAll(file)
	if err != nil {
		fmt.Fprintln(out, err.Error())
		return
	}

	fmt.Fprintln(out, string(b))
}