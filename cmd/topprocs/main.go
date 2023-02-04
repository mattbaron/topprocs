package main

import (
	"fmt"

	"github.com/mattbaron/topprocs/line"
	"github.com/mattbaron/topprocs/procs"
)

func main() {
	filter := procs.NewFilter()
	interestingProcs := procs.FindInteresting(filter)

	for _, proc := range interestingProcs {
		line := line.NewLine("topprocs")
		line.AddTags(proc.Tags())
		line.AddFields(proc.Fields())
		fmt.Println(line.ToString())
	}
}
