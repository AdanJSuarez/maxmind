package app

import (
	"fmt"
	"sync"

	"github.com/AdanJSuarez/maxmind/internal/configuration"
	"github.com/AdanJSuarez/maxmind/internal/geoinfo"
	"github.com/AdanJSuarez/maxmind/internal/logparser"
	"github.com/AdanJSuarez/maxmind/internal/logreader"
	"github.com/AdanJSuarez/maxmind/internal/report"
)

const (
	channelSize     = 1000
	numOfGoroutines = 2
)

type App struct {
	wg        *sync.WaitGroup
	config    configuration.Configuration
	logReader logReader
	geoInfo   geoInfo
	logParser logParser
	report    countryReport
	linesCh   chan string
}

func New(wg *sync.WaitGroup, config configuration.Configuration) (*App, error) {
	app := &App{
		wg:      wg,
		config:  config,
		linesCh: make(chan string, channelSize),
	}
	app.logReader = logreader.New(app.wg, app.config.LogFile(), app.linesCh)
	app.geoInfo = geoinfo.New(app.config.DBfile())

	if err := app.initializeApp(wg); err != nil {
		return nil, err
	}

	return app, nil
}

// Start starts the application. It read lines from the file and populate data for
// the report
func (a *App) Start() {
	a.wg.Add(numOfGoroutines)
	go a.logReader.ReadLinesFromFile()
	go a.populateData()
	a.wg.Wait()

	a.report.Generate()
}

// Close closes the geoInfo database.
func (a *App) Close() {
	if err := a.geoInfo.Close(); err != nil {
		fmt.Printf("error closing the db: %v\n", err)
	}
	if err := a.logReader.Close(); err != nil {
		fmt.Printf("error closing the file: %v\n", err)
	}
}

// initializeApp initializes all the dependencies. It returns an error otherwise.
func (a *App) initializeApp(wg *sync.WaitGroup) error {
	if err := a.setLogParser(); err != nil {
		return err
	}

	if err := a.openLogReader(); err != nil {
		return err
	}

	if err := a.openGeoInfo(); err != nil {
		return err
	}

	a.report = report.New()
	return nil
}

func (a *App) setLogParser() error {
	logParser, err := logparser.New()
	if err != nil {
		fmt.Printf("error on log parser: %v\n", err)
	}

	a.logParser = logParser
	return nil
}

func (a *App) openLogReader() error {
	if err := a.logReader.Open(); err != nil {
		return err
	}
	return nil
}

func (a *App) openGeoInfo() error {
	if err := a.geoInfo.OpenDB(); err != nil {
		fmt.Printf("error open db: %v\n", err)
		return err
	}
	return nil
}

func (a *App) populateData() {
	defer a.wg.Done()

	for line := range a.linesCh {
		lineLog, err := a.logParser.Parse(line)
		if err != nil {
			fmt.Printf("log line error: %s: %v\n", line, err)
			continue
		}

		if a.report.ShouldExclude(lineLog.RequestPath) {
			continue
		}
		record := a.geoInfo.GetIPInfo(lineLog.IP)
		a.report.AddData(record.CountryName, a.report.Subdivision(record.Subdivisions), lineLog.RequestPath)
	}
}
