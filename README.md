# MaxMind

MaxMind log parser that report IP geo-location analysis.

## Requirements

The requirements for this exercise are posted in a [MaxMind public repo](https://github.com/maxmind/dev-hire-homework)

## Considerations

The exercise require to parse a log file that could be bigger than the available memory, so I implement a line by line reader. Also to get the maximum potential to golang
I decided to do as much as possible concurrently.


# Libraries:
https://github.com/xojoc/logparse

https://clavinjune.dev/en/blogs/create-log-parser-using-go/
