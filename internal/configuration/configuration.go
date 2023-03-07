package configuration

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/afero"
)

const (
	dbFileDefault  = "GeoLite2-City.mmdb"
	logFileDefault = "access.log"
)

type Configuration struct {
	DBfile  string
	LogFile string
	fileSys afero.Fs
}

func New() Configuration {
	config := Configuration{}
	config.fileSys = afero.NewOsFs()
	dbFileFlag, logFileFlag := config.flags()
	config.DBfile = dbFileFlag
	config.LogFile = logFileFlag

	return config
}

// CheckConfiguration returns an error if any of the files don't exist.
func (c *Configuration) CheckConfiguration() error {
	var err error
	_, err1 := c.fileSys.Stat(c.DBfile)
	if err1 != nil {
		fmt.Printf("failed on file: %s: %v", c.DBfile, err)
		err = err1
	}
	_, err2 := c.fileSys.Stat(c.LogFile)
	if err2 != nil {
		fmt.Printf("failed on file: %s: %v", c.LogFile, err)
		err = err2
	}
	return err
}

// flags reads and parses the flags from the command line, if any.
// It sets the default values if no flag is passed.
func (c *Configuration) flags() (string, string) {
	dbFileFlag := flag.String("dbFile", dbFileDefault, "dbFile: is the path of the City db")
	logFileFlag := flag.String("logFile", logFileDefault, "logFile: is the path of the log file")
	helpFlag := flag.Bool("help", false, "help: gives information about the flags")
	flag.Parse()

	if *helpFlag {
		fmt.Println("==> Accepted flags:")
		flag.PrintDefaults()
		fmt.Println("==> If there is no flags, it would expect to find the default values in the same folder")
		c.exitPeacefully()
	}
	return *dbFileFlag, *logFileFlag
}

// exitPeacefully exit successfully when help flag is invoked.
func (c *Configuration) exitPeacefully() {
	os.Exit(0)
}
