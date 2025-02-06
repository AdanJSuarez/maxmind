# MaxMind

MaxMind log parser that reports IP geo-location analysis.

## Considerations

The project require **Go 1.18 or higher**. The project comes with `vendor` including all the dependencies to compile it.

To compile the project, if you have installed `make`
```make
make build
```

Otherwise: `go build -o ./bin/maxmind -ldflags "-s -w" ./cmd/maxmind.go`

To run the executable move to `bin/` folder and called `./maxmind`. You can run `./maxmind -help` which explains the flags.
If you don't pass any flag, it assumes the files are located in the same folder with the default names as explained in `help`:

    - `GeoLite2-City.mmdb` for db
    - `access.log` for logs.

`maxmind` expects a [well formatted log file](https://httpd.apache.org/docs/2.4/logs.html#combined). It will not report other than `log errors` if the file is not well formatted.

If any record in the log file is corrupted or malformed, it will report a `log line error`.

## Unit Test

It requires [Mockery](https://vektra.github.io/mockery/installation/) to generate the mocks, needed for the unit tests to run. There are different ways to generate the mocks explained in the Mockery documentation.

- If you install Mockery with `go install github.com/vektra/mockery/v2@v2.20.0` after the installation you just need to either `go generate ./...` or `make mock`

- If you want to use the Docker image you **must** need to include the flag `--inpackage`.

Note: Unfortunately the mock files are included in the coverage so they reduce the actual coverage of the unit test suites which in most cases is 100%. For a visual and more accurate coverage analysis, you could type `make cover FOLDER="path_to_folder"`
