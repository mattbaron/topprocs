package procs

import (
	"fmt"

	"github.com/mattbaron/topprocs/line"
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
	MemoryInfo    process.MemoryInfoStat
	NumThreads    int32
	User          string
}

func NewProc(pid int32) *Proc {

	process, err := process.NewProcess(pid)

	if err != nil {
		return nil
	}

	proc := Proc{Process: process}
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

	memoryInfo, err := proc.Process.MemoryInfo()
	if err == nil {
		proc.MemoryInfo = *memoryInfo
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
