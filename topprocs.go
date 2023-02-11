package topprocs

//go:embed sample.conf
var sampleConfig string

import (
	"fmt"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

type TopProcs struct {

}

func (topProcs *TopProcs) Init() error {
	return nil
}

func (topProcs *TopProcs) Gather(acc telegraf.Accumulator) error {
	return nil
}

func init() {
	inputs.Add("topprocs", func() telegraf.Input {
		&TopProcs{}
	})
}
