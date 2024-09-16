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

// go run grepdirrec.go match testdata/file3.txt
// go run grepdirrec.go match testdata

func main() {
	GrepDirRec(os.Args[1], os.Args[2], os.Stdout)
}

func GrepDirRec(match string, recPath string, out io.Writer) {
	info, err := os.Stat(recPath)
	if err != nil {
		fmt.Fprintln(out, err.Error())
		return
	}

	if info.IsDir() {
		filepath.WalkDir(recPath, func(path string, d fs.DirEntry, err error) error {
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
	} else {
		go grepAndPrint(match, recPath, out)
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