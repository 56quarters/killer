// Killer - Repeatedly try to kill a process
//
// Copyright 2017 TSH Labs
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"syscall"
	"time"
)

const (
	DEFAULT_INTERVAL = 1
	DEFAULT_TIMEOUT  = 30
)

func killNicely(p *os.Process, interval int, timeout int) (bool, error) {
	elapsed := 0

	for {
		res := p.Signal(syscall.SIGTERM)
		if res == syscall.ESRCH {
			return true, nil
		} else if res == syscall.EPERM {
			return false, res
		} else if res == syscall.EINVAL {
			return false, res
		}

		if elapsed >= timeout {
			return false, nil
		}

		time.Sleep(time.Duration(interval * 1000000000))
		elapsed += interval
	}
}

func killNotSoNicely(p *os.Process) error {
	res := p.Signal(syscall.SIGKILL)
	// Successfully sent the signal or it's already stopped
	if res == nil || res == syscall.ESRCH {
		return nil
	}

	return res
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
