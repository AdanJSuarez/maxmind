package logreader

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type LogReader struct {
	linesCh chan string
}

func New(linesCh chan string) *LogReader {
	return &LogReader{
		linesCh: linesCh,
	}
}

func (fr *LogReader) ReadLinesFromFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Panicf("error opening the file: %s: %v", path, err)
	}
	defer file.Close()

	fr.sendLinesToLinesCh(file)
}

func (fr *LogReader) sendLinesToLinesCh(file *os.File) {
	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fmt.Println("on scanner ==>", fileScanner.Text())
		fr.linesCh <- fileScanner.Text()
	}
}
