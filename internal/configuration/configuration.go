package configuration

import (
	"flag"
	"fmt"
	"os"
)

const (
	dbFile  = "GeoLite2-City.mmdb"
	logFile = "access.log"
)

type Configuration struct {
	DBfile  string
	LogFile string
}

func New() Configuration {
	dbFileFlag := flag.String("dbFile", dbFile, "dbFile: is the path of the City db")
	logFileFlag := flag.String("logFile", logFile, "logFile: is the path of the log file")
	helpFlag := flag.Bool("help", false, "help: gives information about the flags")
	flag.Parse()

	if *helpFlag {
		fmt.Println("==> Accepted flags:")
		flag.PrintDefaults()
		fmt.Println("==> If there is no flags, it would expect to find the default values in the same folder")
		exitPeacefully()
	}
	return Configuration{
		DBfile:  *dbFileFlag,
		LogFile: *logFileFlag,
	}
}

func exitPeacefully() {
	os.Exit(0)
}
