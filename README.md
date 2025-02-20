# last-minute

Tail last minute of a log file. Usage:

    $ tail -F main.log | lastmin --seconds 5

Every time tail feeds another line, terminal screen is cleared, and most recent
n seconds (or minutes) of logs are printed.

[![Demo](https://asciinema.org/a/703573.svg)](https://asciinema.org/a/703573)

# Usage

    Usage: lastmin --seconds=INT --minutes=INT [flags]
    
    Flags:
      -h, --help                 Show context-sensitive help.
      -s, --seconds=INT
      -m, --minutes=INT
      -t, --show-tick            Show current time.
      -r, --refresh-seconds=2    How often output is refreshed.

# Install

    go install github.com/denarced/last-minute@latest
