# MaxMind

MaxMind log parser that report IP geo-location analysis.

## Homework requirements

The requirements for this exercise are posted in a [MaxMind public repo](https://github.com/maxmind/dev-hire-homework)

## Considerations

The project needs Go 1.18+

To compile the project, if you have installed `make`
```make
make build
```

To run the executable called `maxmind` you can run `maxmind -help` which explain the flags.
If you don't pass any flag, it assumes the files are located in the same folder with the default names as explained in `help`, `GeoLite2-City.mmdb` for db and `access.log` for logs.

`maxmind` expects a well formatted log file. It will not report other than `log errors` if the file is not well formatted.


## Unit Test

It requires Mockery for generating the mocks, needed for the unit tests to run. See instruction on Mockery to how...
TODO: FINISH Readme.md


# Libraries:
https://github.com/xojoc/logparse

https://clavinjune.dev/en/blogs/create-log-parser-using-go/
