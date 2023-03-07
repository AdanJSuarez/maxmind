package logreader

import (
	"bufio"
	"fmt"
	"sync"

	"github.com/spf13/afero"
)

type LogReader struct {
	linesCh  chan string
	fileSys  afero.Fs
	file     afero.File
	filePath string
}

func New(filePath string, linesCh chan string) (*LogReader, error) {
	logReader := &LogReader{}
	afero := afero.NewOsFs()
	logReader.fileSys = afero
	logReader.linesCh = linesCh
	logReader.filePath = filePath

	file, err := logReader.fileSys.Open(logReader.filePath)
	if err != nil {
		return nil, err
	}

	logReader.file = file
	return logReader, nil
}

/*
ReadLinesFromFile returns nil and sends the lines read through the channel.
It close the channel to sync with the receiver indicating the end of the file.
*/
func (lr *LogReader) ReadLinesFromFile(wg *sync.WaitGroup) {
	defer wg.Done()
	defer lr.file.Close()
	fmt.Printf("==> Reading lines of file: %s\n", lr.filePath)

	lr.sendLinesToLinesCh()
}

// sendLinesToLinesCh reads line by line and sends them to linesCh.
// It close the channel to sync with the receiver when it finished.
func (lr *LogReader) sendLinesToLinesCh() {
	defer close(lr.linesCh)

	fileScanner := bufio.NewScanner(lr.file)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		lr.linesCh <- fileScanner.Text()
	}
}
