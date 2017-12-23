package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

const DEFAULT_INTERVAL = 1
const DEFAULT_TIMEOUT = 30

func killNicely(p *os.Process, interval int, timeout int) (bool, error) {
	return true, nil
}

func killNotSoNicely(p *os.Process) error {
	return nil
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] [pid]\n", os.Args[0])
		flag.PrintDefaults()
	}

	interval := flag.Int("interval", DEFAULT_INTERVAL, "How long to wait between attempts to stop a process in seconds")
	timeout := flag.Int("timeout", DEFAULT_TIMEOUT, "How long to wait total when trying to stop a process in seconds")
	kill9 := flag.Bool("use-kill", true, "Should SIGKILL be used as a last resort when stopping a process")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Printf("PID is required\n")
		os.Exit(1)
	}

	pid, err := strconv.ParseInt(flag.Args()[0], 10, 32)
	if err != nil {
		log.Fatal(err)
	}

	p, err := os.FindProcess(int(pid))
	if err != nil {
		log.Fatal(err)
	}

	stopped, err := killNicely(p, *interval, *timeout)
	if err != nil {
		log.Fatal(err)
	}

	if !stopped && *kill9 {
		killNotSoNicely(p)
	}
}
