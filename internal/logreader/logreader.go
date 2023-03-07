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

func New(filePath string, linesCh chan string) *LogReader {
	return &LogReader{
		fileSys:  afero.NewOsFs(),
		linesCh:  linesCh,
		filePath: filePath,
	}
}

func (lr *LogReader) Open() error {
	file, err := lr.fileSys.Open(lr.filePath)
	if err != nil {
		return err
	}

	lr.file = file
	return nil
}

func (lr *LogReader) Close() error {
	return lr.file.Close()
}

/*
ReadLinesFromFile returns nil and sends the lines read through the channel.
It close the channel to sync with the receiver indicating the end of the file.
*/
func (lr *LogReader) ReadLinesFromFile(wg *sync.WaitGroup) {
	defer wg.Done()
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
