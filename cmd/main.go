package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mattbaron/topprocs/influx"
	"github.com/mattbaron/topprocs/procs"
	"github.com/mattbaron/topprocs/version"
)

var progname = filepath.Base(os.Args[0])

func main() {
	filter := procs.NewFilter()

	versionFlag := flag.Bool("version", false, "help message for flag n")
	debugFlag := flag.Bool("debug", false, "Debugging")

	flag.Float64Var(&filter.CPUUsage, "cpu_usage", 1.0, "CPU filter")
	flag.Float64Var(&filter.MemoryUsage, "memory_usage", 1.0, "Memory filter")
	flag.Int64Var(&filter.NumThreads, "num_threads", 50, "Thread filter")
	flag.Parse()

	if *versionFlag {
		fmt.Println(version.Version)
		os.Exit(0)
	}

	interestingProcs := procs.FindInteresting(filter, *debugFlag)

	for _, proc := range interestingProcs {
		if proc.Name == progname {
			continue
		}

		line := influx.NewLine("topprocs")
		line.AddTags(proc.Tags())
		line.AddFields(proc.Fields())
		fmt.Println(line.ToString())
	}
}
