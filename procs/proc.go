package procs

import (
	"fmt"

	"github.com/mattbaron/topprocs/line"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/process"
)

type Proc struct {
	p             *process.Process
	Pid           int32
	Name          string
	CPUPercent    float64
	MemoryPercent float32
	CPUTime       cpu.TimesStat
	MemoryInfo    process.MemoryInfoStat
	NumThreads    int32
	User          string
}

func NewProc(pid int32) *Proc {

	process, err := process.NewProcess(pid)

	if err != nil {
		return nil
	}

	proc := Proc{p: process}
	err = proc.GatherMetrics()

	if err != nil {
		return nil
	} else {
		return &proc
	}
}

func (proc *Proc) GatherMetrics() error {
	proc.Pid = proc.p.Pid

	cpuPercent, err := proc.p.CPUPercent()
	if err == nil {
		proc.CPUPercent = cpuPercent
	} else {
		return err
	}

	memPercent, err := proc.p.MemoryPercent()
	if err == nil {
		proc.MemoryPercent = memPercent
	} else {
		return err
	}

	numThreads, err := proc.p.NumThreads()
	if err == nil {
		proc.NumThreads = numThreads
	} else {
		return err
	}

	name, err := proc.p.Name()
	if err == nil {
		proc.Name = name
	} else {
		return err
	}

	user, err := proc.p.Username()
	if err == nil {
		proc.User = user
	} else {
		return err
	}

	memoryInfo, err := proc.p.MemoryInfo()
	if err == nil {
		proc.MemoryInfo = *memoryInfo
	} else {
		return err
	}

	cpuTime, err := proc.p.Times()
	if err == nil {
		proc.CPUTime = *cpuTime
	} else {
		return err
	}

	return nil
}

func (proc Proc) Tags() line.Tags {
	tags := line.Tags{
		"name": proc.Name,
		"pid":  fmt.Sprint(proc.Pid),
		"user": proc.User,
	}

	return tags
}

func (proc Proc) Fields() line.Fields {
	fields := line.Fields{
		"cpu_usage":       proc.CPUPercent,
		"memory_usage":    proc.MemoryPercent,
		"memory_rss":      proc.MemoryInfo.RSS,
		"memory_vms":      proc.MemoryInfo.VMS,
		"memory_swap":     proc.MemoryInfo.Swap,
		"num_threads":     proc.NumThreads,
		"cpu_time_iowait": proc.CPUTime.Iowait,
		"cpu_time_user":   proc.CPUTime.User,
		"cpu_time_system": proc.CPUTime.System,
	}

	return fields
}

func (proc *Proc) ToString() string {
	return fmt.Sprintf("name=%s, cpu=%f, mem=%f, t=%d", proc.Name, proc.CPUPercent, proc.MemoryPercent, proc.NumThreads)
}

func AllPids() []int32 {
	pids, err := process.Pids()

	if err != nil {
		return make([]int32, 0)
	}

	return pids
}

func FindInteresting(filter Filter) []*Proc {
	procs := make([]*Proc, 0)

	for _, pid := range AllPids() {
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
