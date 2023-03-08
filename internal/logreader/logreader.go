package logreader

import (
	"bufio"
	"fmt"
	"sync"

	"github.com/spf13/afero"
)

type LogReader struct {
	wg       *sync.WaitGroup
	linesCh  chan string
	fileSys  afero.Fs
	file     afero.File
	filePath string
}

// New returns a initialized instance of LogReader.
func New(wg *sync.WaitGroup, filePath string, linesCh chan string) *LogReader {
	return &LogReader{
		wg:       wg,
		fileSys:  afero.NewOsFs(),
		linesCh:  linesCh,
		filePath: filePath,
	}
}

// Open opens the log file. It returns an error otherwise.
func (lr *LogReader) Open() error {
	file, err := lr.fileSys.Open(lr.filePath)
	if err != nil {
		return err
	}

	lr.file = file
	return nil
}

// Close closes the log file. It returns an error otherwise.
func (lr *LogReader) Close() error {
	if lr.file != nil {
		return lr.file.Close()
	}
	return nil
}

/*
ReadLinesFromFile sends log lines read through the channel.
It close the channel to sync with the receiver indicating the end of the file.
*/
func (lr *LogReader) ReadLinesFromFile() {
	defer lr.wg.Done()
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
