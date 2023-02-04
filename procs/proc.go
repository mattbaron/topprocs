package procs

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/process"
)

type Proc struct {
	*process.Process
	Pid           int32
	Name          string
	CPUPercent    float64
	MemoryPercent float32
	CPUTime       cpu.TimesStat
	Memory        process.MemoryInfoExStat
	NumThreads    int32
	User          string
}

func NewProc(p *process.Process) *Proc {
	proc := Proc{Process: p}
	proc.GatherMetrics()
	return &proc
}

func (proc *Proc) GatherMetrics() error {
	proc.Pid = proc.Process.Pid

	cpuPercent, err := proc.Process.CPUPercent()
	if err == nil {
		proc.CPUPercent = cpuPercent
	} else {
		return err
	}

	memPercent, err := proc.Process.MemoryPercent()
	if err == nil {
		proc.MemoryPercent = memPercent
	} else {
		return err
	}

	numThreads, err := proc.Process.NumThreads()
	if err == nil {
		proc.NumThreads = numThreads
	} else {
		return err
	}

	name, err := proc.Process.Name()
	if err == nil {
		proc.Name = name
	} else {
		return err
	}

	user, err := proc.Process.Username()
	if err == nil {
		proc.User = user
	} else {
		return err
	}

	cpuTime, err := proc.Process.Times()
	if err == nil {
		proc.CPUTime = *cpuTime
	} else {
		return err
	}

	return nil
}

func (proc *Proc) Filter(filter Filter) {

}
