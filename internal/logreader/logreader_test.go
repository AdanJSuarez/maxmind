package logreader

import (
	"sync"
	"testing"
	"time"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/suite"
)

const (
	fileNameTest = "fakeFile.log"
	fakeLog      = "183.60.212.148 - - [26/Aug/2014:06:26:39 -0600]"
)

var (
	readerTest  *LogReader
	wgTest      *sync.WaitGroup
	linesChTest chan string
)

type TSLogReader struct{ suite.Suite }

func TestRunTSLogReader(t *testing.T) {
	suite.Run(t, new(TSLogReader))
}

func (ts *TSLogReader) BeforeTest(_, _ string) {
	wgTest = &sync.WaitGroup{}
	linesChTest = make(chan string, 10)
	readerTest = New(wgTest, fileNameTest, linesChTest)
	readerTest.fileSys = afero.NewMemMapFs()
	afero.WriteFile(readerTest.fileSys, fileNameTest, []byte(fakeLog), 0644)
}

func (ts *TSLogReader) TestReadLinesFromFileForValidFile() {
	err := readerTest.Open()
	ts.NoError(err)

}

func (ts *TSLogReader) TestCloseWithoutOpenIt() {
	err := readerTest.Close()
	ts.NoError(err)
}

func (ts *TSLogReader) TestCloseAfterOpenIt() {
	err1 := readerTest.Open()
	ts.NoError(err1)
	err2 := readerTest.Close()
	ts.NoError(err2)
}

func (ts *TSLogReader) TestReadLinesFromFileForInvalidFile() {
	readerTest := New(wgTest, "fakeFile.log", linesChTest)
	ts.NotNil(readerTest)
	err := readerTest.Open()
	ts.ErrorContains(err, "no such file or directory")

}

func (ts *TSLogReader) TestReadLineFromFile() {
	wgTest.Add(1)
	condition := func() bool {
		log := <-readerTest.linesCh
		return len(log) > 0
	}
	err := readerTest.Open()
	ts.NoError(err)

	go readerTest.ReadLinesFromFile()
	ts.Eventually(condition, 5*time.Second, 200*time.Microsecond)
}
