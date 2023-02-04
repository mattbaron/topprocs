package procs

import (
	"github.com/shirou/gopsutil/process"
)

type ProcessList struct {
	procs map[int32]*process.Process
}

func NewProcessList() *ProcessList {
	return &ProcessList{
		procs: make(map[int32]*process.Process, 0),
	}
}

func (processList *ProcessList) AddProcess(p *process.Process) {
	processList.procs[p.Pid] = p
}

func FindAll() ([]int32, error) {
	procs, err := process.Pids()

	if err != nil {
		return nil, err
	}

	return procs, nil
}
