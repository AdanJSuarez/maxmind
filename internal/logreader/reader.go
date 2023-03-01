package logreader

import (
	"bufio"
	"fmt"
	"log"

	"github.com/spf13/afero"
)

type LogReader struct {
	linesCh chan string
	fileSys afero.Fs
}

func New(linesCh chan string) *LogReader {
	afero := afero.NewOsFs()

	return &LogReader{
		linesCh: linesCh,
		fileSys: afero,
	}
}

func (fr *LogReader) ReadLinesFromFile(path string) {
	file, err := fr.fileSys.Open(path)
	if err != nil {
		log.Panicf("error opening the file: %s: %v", path, err)
	}
	defer file.Close()

	fr.sendLinesToLinesCh(file)
}

func (fr *LogReader) sendLinesToLinesCh(file afero.File) {
	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fmt.Println("on scanner ==>", fileScanner.Text())
		fr.linesCh <- fileScanner.Text()
	}
}
