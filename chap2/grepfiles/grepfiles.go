package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// go run grepfiles.go match ./testdata/file1.txt ./testdata/file2.txt ./testdata/file3.txt

func main() {
	GrepFiles(os.Args[1], os.Args[2:], os.Stdout)
}

func GrepFiles(match string, fileNames []string, out io.Writer) {
	for _, fileName := range fileNames {
		go grepAndPrint(match, fileName, out)
	}

	time.Sleep(3 * time.Second)
}

func grepAndPrint(match, fileName string, out io.Writer) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintln(out, err.Error())
		return
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), match) {
			fmt.Fprintln(out, scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(out, err.Error())
	}
}