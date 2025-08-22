
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"net-logger/internal/core"
	"net-logger/internal/adapters"
	"net-logger/internal/app"
)

func main() {
	var (
		untilStr string
		forStr string
	)
	flag.StringVar(&untilStr, "until", "", "Run until this datetime (RFC3339)")
	flag.StringVar(&forStr, "for", "", "Run for this duration (e.g. 10s, 1m)")
	flag.Parse()
	targets := flag.Args()

	if len(targets) == 0 {
		fmt.Println("No targets specified. Exiting.")
		os.Exit(1)
	}

	var (
		ctx context.Context
		cancel context.CancelFunc
	)

	if untilStr != "" {
		until, err := time.Parse(time.RFC3339, untilStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid --until: %v\n", err)
			os.Exit(1)
		}
		dur := time.Until(until)
		ctx, cancel = context.WithTimeout(context.Background(), dur)
	} else if forStr != "" {
		dur, err := time.ParseDuration(forStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid --for: %v\n", err)
			os.Exit(1)
		}
		ctx, cancel = context.WithTimeout(context.Background(), dur)
	} else {
		ctx, cancel = context.WithCancel(context.Background())
	}
	defer cancel()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		<-c
		fmt.Println("\nReceived interrupt, exiting...")
		cancel()
	}()

	results := &core.SafeResults{}
	runner := &app.Runner{
		PingerICMP: &adapters.ICMPPinger{},
		PingerTCP: &adapters.TCPPinger{},
		Results: results,
	}

	wg := sync.WaitGroup{}
	for _, t := range targets {
		t := t
		connType := app.DetectConnType(t)
		wg.Add(1)
		go func() {
			defer wg.Done()
			runner.ProbeLoop(ctx, t, connType)
		}()
	}
	wg.Wait()

	reporter := &adapters.MarkdownReporter{}
	reporter.Report(results.All())
}

