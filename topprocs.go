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
	Log         telegraf.Logger
	CPUUsage    float64 `toml:"cpu_usage"`
	MemoryUsage float64 `toml:"memory_usage"`
	NumThreads  int64   `toml:"num_threads"`
	MemoryVMS   uint64  `toml:"memory_vms"`
	Filter      procs.Filter
}

func (*TopProcs) SampleConfig() string {
	return sampleConfig
}

func (topProcs *TopProcs) Gather(acc telegraf.Accumulator) error {
	topProcs.Log.Infof("Gather() Filter: %v\n", topProcs.Filter)
	interestingProcs := procs.FindInteresting(topProcs.Filter, false)
	for _, proc := range interestingProcs {
		acc.AddFields("topprocs", proc.Fields(), proc.Tags())
	}

	return nil
}

func (topProcs *TopProcs) Init() error {
	topProcs.Log.Info("Init()")
	topProcs.Filter = procs.Filter{
		CPUUsage:    topProcs.CPUUsage,
		MemoryUsage: topProcs.MemoryUsage,
		NumThreads:  topProcs.NumThreads,
		MemoryVMS:   topProcs.MemoryVMS,
	}
	return nil
}

func init() {
	inputs.Add("topprocs", func() telegraf.Input {
		return &TopProcs{
			CPUUsage:    1.0,
			MemoryUsage: 1.0,
			NumThreads:  50,
			MemoryVMS:   1073741824,
		}
	})
}
