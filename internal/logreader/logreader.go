package logreader

import (
	"bufio"
	"fmt"

	"github.com/spf13/afero"
)

type LogReader struct {
	linesCh  chan string
	fileSys  afero.Fs
	filePath string
}

func New(filePath string, linesCh chan string) (*LogReader, error) {
	logReader := &LogReader{}
	afero := afero.NewOsFs()
	logReader.fileSys = afero
	logReader.linesCh = linesCh
	logReader.filePath = filePath

	if err := logReader.checkFile(); err != nil {
		return nil, err
	}

	return logReader, nil
}

/*
ReadLinesFromFile returns nil and sends through the channel the lines read.
It close the channel to sync with the receiver indicating the end of the file.
It returns an error if cannot open the file from the path.
*/
func (lr *LogReader) ReadLinesFromFile() error {
	file, err := lr.fileSys.Open(lr.filePath)
	if err != nil {
		return fmt.Errorf("error opening the file: %s: %v", lr.filePath, err)
	}
	defer file.Close()

	lr.sendLinesToLinesCh(file)
	return nil
}

// sendLinesToLinesCh reads line by line and sends them to linesCh.
// It close the channel to sync with the receiver when it finished.
func (lr *LogReader) sendLinesToLinesCh(file afero.File) {
	defer close(lr.linesCh)

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		lr.linesCh <- fileScanner.Text()
	}
}

func (lr *LogReader) checkFile() error {
	_, err := lr.fileSys.Stat(lr.filePath)
	if err != nil {
		return fmt.Errorf("error on file: %s: %v", lr.filePath, err)
	}
	return err
}
