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
	channelSize = 1000
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

	if err := app.initializeApp(wg); err != nil {
		return nil, err
	}

	return app, nil
}

// Start starts the application. It read lines from the file and populate data for
// the report
func (a *App) Start() {
	a.wg.Add(2)
	go a.logReader.ReadLinesFromFile(a.wg)
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

	if err := a.setLogReader(a.linesCh); err != nil {
		return err
	}

	if err := a.setGeoInfo(); err != nil {
		return err
	}

	a.report = report.New()
	return nil
}

func (a *App) setGeoInfo() error {
	a.geoInfo = geoinfo.New(a.config.DBfile)
	if err := a.geoInfo.OpenDB(); err != nil {
		fmt.Printf("error open db: %v\n", err)
		return err
	}

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

func (a *App) setLogReader(linesCh chan string) error {
	logReader, err := logreader.New(a.config.LogFile, linesCh)
	if err != nil {
		return err
	}

	a.logReader = logReader
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
