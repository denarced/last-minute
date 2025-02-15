// Package main implements the main CLI.
package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/alecthomas/kong"
	"github.com/denarced/last-minute/lib/lastmin"
	"github.com/denarced/last-minute/shared"
	"github.com/inancgumus/screen"
)

var CLI struct {
	Seconds        int  `short:"s" xor:"Minutes,Seconds" required:"true"`
	Minutes        int  `short:"m" xor:"Minutes,Seconds" required:"true"`
	ShowTick       bool `short:"t" default:"false" help:"Show current time."`
	RefreshSeconds int  `short:"r" default:"2" help:"How often output is refreshed."`
}

func main() {
	shared.InitLogging()
	kong.Parse(&CLI)
	shared.Logger.Info("Start.", "CLI arguments", os.Args)

	seconds := CLI.Minutes * 60
	if code := deriveErrorExitCode(); code > 0 {
		os.Exit(code)
	}
	if CLI.Seconds > 0 {
		seconds = CLI.Seconds
	}

	feed := make(chan string)
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			feed <- scanner.Text()
		}
	}()

	screen.Clear()
	lines := []lastmin.DatedLine{}
	now := time.Now()
mainLoop:
	for {
		select {
		case <-time.After(time.Duration(CLI.RefreshSeconds) * time.Second):
		case line, ok := <-feed:
			if !ok {
				break mainLoop
			}
			date, err := lastmin.ParseDate(line)
			if err == nil {
				lines = append(lines, lastmin.DatedLine{Date: date, Line: line})
			} else {
				shared.Logger.Warn("No date found, ignoring.", "line", line)
			}
		}
		lines = lastmin.FilterLines(lines, now, time.Now(), seconds)
		screen.Clear()
		if CLI.ShowTick {
			fmt.Println("time", time.Now())
		}
		for _, each := range lines {
			fmt.Println(each.Line)
		}
	}
	shared.Logger.Info("Done.")
}

func deriveErrorExitCode() (code int) {
	if CLI.Seconds <= 0 && CLI.Minutes <= 0 {
		fmt.Fprintf(os.Stderr, "Either --minutes or --seconds must be >0.\n")
		code = 2
	}
	if CLI.RefreshSeconds < 1 {
		fmt.Fprintf(
			os.Stderr,
			"Illegal value for refresh seconds: %d. Must be >0.\n",
			CLI.RefreshSeconds,
		)
		if code == 0 {
			code = 4
		}
	}
	return
}
