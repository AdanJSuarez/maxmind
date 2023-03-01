package internal

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type FileReader struct {
	linesCh chan string
}

func New(linesCh chan string) *FileReader {
	return &FileReader{
		linesCh: linesCh,
	}
}

func (fr *FileReader) ReadLinesFromFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Panicf("error opening the file: %s: %v", path, err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fmt.Println("on scanner ==>", fileScanner.Text())
		fr.linesCh <- fileScanner.Text()
	}
}
