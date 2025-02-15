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
	dbFileFlag     = "dbfile"
	logFileFlag    = "logfile"
	helpFlag       = "help"
	dbfileMessage  = "dbfile: is the path of the City db"
	logfileMessage = "logfile: is the path of the log file"
	helpMessage    = "help: gives information about the flags"
)

type Configuration struct {
	dbFile  string
	logFile string
	fileSys afero.Fs
}

// New returns an initialized instance of Configuration.
func New() Configuration {
	config := Configuration{}
	config.fileSys = afero.NewOsFs()
	dbFileFlag, logFileFlag := config.flags()
	config.dbFile = dbFileFlag
	config.logFile = logFileFlag

	return config
}

// DBFile db file path.
func (c *Configuration) DBfile() string {
	return c.dbFile
}

// LogFile log file path.
func (c *Configuration) LogFile() string {
	return c.logFile
}

// CheckConfiguration returns an error if any of the files don't exist.
func (c *Configuration) CheckConfiguration() error {
	var err error

	_, err1 := c.fileSys.Stat(c.dbFile)
	if err1 != nil {
		fmt.Printf("failed on file: %s: %v\n", c.dbFile, err1)
		err = err1
	}

	_, err2 := c.fileSys.Stat(c.logFile)
	if err2 != nil {
		fmt.Printf("failed on file: %s: %v\n", c.logFile, err2)
		err = err2
	}
	return err
}

// flags reads and parses the flags from the command line, if any.
// It sets the default values if no flag is passed.
func (c *Configuration) flags() (string, string) {
	dbFileFlag := flag.String(dbFileFlag, dbFileDefault, dbfileMessage)
	logFileFlag := flag.String(logFileFlag, logFileDefault, logfileMessage)
	helpFlag := flag.Bool(helpFlag, false, helpMessage)
	flag.Parse()

	if *helpFlag {
		fmt.Println("==> Accepted flags:")
		flag.PrintDefaults()
		fmt.Println("==> If there is no flags, it would expect to find the default values in the same folder")
		c.exitPeacefully()
	}

	return *dbFileFlag, *logFileFlag
}

// exitPeacefully exit successfully when "help" flag is invoked.
func (c *Configuration) exitPeacefully() {
	os.Exit(0)
}
