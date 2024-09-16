package main

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// go run grepdir.go match testdata

func main() {
	GrepDir(os.Args[1], os.Args[2], os.Stdout)
}

func GrepDir(match string, dirName string, out io.Writer) {
	filepath.WalkDir(dirName, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Fprintln(out, err.Error())
			return err
		}

		if d.IsDir() {
			return nil
		}

		go grepAndPrint(match, path, out)
		return nil
	})

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