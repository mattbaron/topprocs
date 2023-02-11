package topprocs

import (
	_ "embed"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
	"github.com/mattbaron/topprocs/procs"
)

//go:embed sample.conf
var sampleConfig string

type TopProcs struct {
	Filter        procs.Filter
	CPUPercent    float64 `toml:"cpu_usage"`
	MemoryPercent float64 `toml:"memory_usage"`
	Threads       int64   `toml:"num_threads"`
	MemoryVMS     uint64  `toml:"memory_vms"`
}

func NewTopProcs() *TopProcs {
	topProcs := TopProcs{}
	topProcs.Filter = procs.Filter{
		CPUPercent:    topProcs.CPUPercent,
		MemoryPercent: topProcs.MemoryPercent,
		Threads:       topProcs.Threads,
		MemoryVMS:     topProcs.MemoryVMS,
	}
	return &topProcs
}

func (*TopProcs) SampleConfig() string {
	return sampleConfig
}

func (topProcs *TopProcs) Gather(acc telegraf.Accumulator) error {
	interestingProcs := procs.FindInteresting(topProcs.Filter, false)
	for _, proc := range interestingProcs {
		acc.AddFields("topprocs", proc.Fields(), proc.Tags())
	}

	return nil
}

func init() {
	inputs.Add("topprocs", func() telegraf.Input {
		return &TopProcs{}
	})
}
