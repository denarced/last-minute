# last-minute

Tail last minute of a log file. Usage:

    $ tail -F main.log | lastmin --seconds 5

Every time tail feeds another line, terminal screen is cleared, and most recent
n seconds (or minutes) of logs are printed.

[![Demo](https://asciinema.org/a/703573.svg)](https://asciinema.org/a/703573)
