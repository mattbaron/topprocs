package procs

import (
	"github.com/shirou/gopsutil/process"
)

type ProcessList struct {
	procs map[int32]*Proc
}

func NewProcessList() *ProcessList {
	return &ProcessList{
		procs: make(map[int32]*Proc, 0),
	}
}

func (processList *ProcessList) AddProcess(p *Proc) {
	processList.procs[p.Pid] = p
}

func FindAll() []int32 {
	procs, err := process.Pids()

	if err != nil {
		return make([]int32, 0)
	}

	return procs
}

func FindInteresting(filter Filter) []*Proc {
	procs := make([]*Proc, 0)

	for _, pid := range FindAll() {
		proc := NewProc(pid)

		if proc == nil {
			continue
		}

		if filter.Match(*proc) {
			procs = append(procs, proc)
		}
	}

	return procs
}
