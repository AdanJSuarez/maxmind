package logreader

import (
	"bufio"
	"fmt"

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

/*
ReadLinesFromFile returns nil and sends through the channel the lines read.
It close the channel to sync with the receiver indicating the end of the file.
It returns an error if cannot open the file from the path.
*/
func (fr *LogReader) ReadLinesFromFile(path string) error {
	file, err := fr.fileSys.Open(path)
	if err != nil {
		return fmt.Errorf("error opening the file: %s: %v", path, err)
	}
	defer file.Close()

	fr.sendLinesToLinesCh(file)
	return nil
}

func (fr *LogReader) sendLinesToLinesCh(file afero.File) {
	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		// fmt.Println("on scanner ==>", fileScanner.Text())
		fr.linesCh <- fileScanner.Text()
	}
	close(fr.linesCh)
}
