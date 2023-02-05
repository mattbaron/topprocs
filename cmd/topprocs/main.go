package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mattbaron/topprocs/line"
	"github.com/mattbaron/topprocs/procs"
)

var progname = filepath.Base(os.Args[0])

func main() {
	filter := procs.NewFilter()

	flag.Float64Var(&filter.CPUPercent, "cpu-percent", 1.0, "CPU filter")
	flag.Float64Var(&filter.MemoryPercent, "mem-percent", 1.0, "Memory filter")
	flag.Int64Var(&filter.Threads, "threads", 50, "Thread filter")
	flag.Parse()

	interestingProcs := procs.FindInteresting(filter)

	for _, proc := range interestingProcs {
		if proc.Name == progname {
			continue
		}

		line := line.NewLine("topprocs")
		line.AddTags(proc.Tags())
		line.AddFields(proc.Fields())
		fmt.Println(line.ToString())
	}
}
