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

	flag.Float64Var(&filter.CPUUsage, "cpu_usage", 1.0, "Minimum CPU usage (percent)")
	flag.Float64Var(&filter.MemoryUsage, "memory_usage", 1.0, "Minimum memory usage (percent)")
	flag.Int64Var(&filter.NumThreads, "num_threads", 50, "Minimum number of threads")
	flag.Int64Var(&filter.AgeMS, "age", 2000, "Minimum process age in milliseconds")
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
