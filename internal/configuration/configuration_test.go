package configuration

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/suite"
)

const (
	accessFileTest = "access.log"
	dbFileTest     = "GeoLite2-City.mmdb"
	fakeLog        = "183.60.212.148 - - [26/Aug/2014:06:26:39 -0600]"
)

var configTest Configuration

type TSConfiguration struct{ suite.Suite }

func TestRunTSConfiguration(t *testing.T) {
	suite.Run(t, new(TSConfiguration))
}

func (ts *TSConfiguration) BeforeTest(_, _ string) {
	configTest = Configuration{
		dbFile:  dbFileTest,
		logFile: accessFileTest,
		fileSys: afero.NewMemMapFs(),
	}
}

func (ts *TSConfiguration) TestInitialization() {
	configTest = New()
	ts.NotNil(configTest)
	ts.Equal("GeoLite2-City.mmdb", configTest.DBfile())
	ts.Equal("access.log", configTest.LogFile())
}

func (ts *TSConfiguration) TestCheckConfigurationOnError() {
	err := configTest.CheckConfiguration()
	ts.Error(err)
}

func (ts *TSConfiguration) TestCheckConfiguration() {
	afero.WriteFile(configTest.fileSys, accessFileTest, []byte(fakeLog), 0644)
	afero.WriteFile(configTest.fileSys, dbFileTest, []byte("fakeDBData"), 0644)

	err := configTest.CheckConfiguration()
	ts.NoError(err)
}
