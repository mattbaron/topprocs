package procs

import (
	"fmt"
	"time"

	"github.com/mattbaron/topprocs/influx"
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
	Errors        []error
	AgeMS         int64
}

func NewProc(pid int32) *Proc {

	process, err := process.NewProcess(pid)

	if err != nil {
		return nil
	}

	proc := Proc{
		p:      process,
		Errors: make([]error, 0),
	}

	err = proc.GatherMetrics()

	if err != nil {
		proc.Errors = append(proc.Errors, err)
	}

	return &proc
}

func (proc *Proc) GatherMetrics() error {
	proc.Pid = proc.p.Pid

	createTime, err := proc.p.CreateTime()
	if err == nil {
		proc.AgeMS = time.Now().UnixMilli() - createTime
	} else {
		return err
	}

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
		uids, err := proc.p.Uids()
		if err == nil && len(uids) > 0 {
			proc.User = fmt.Sprint(uids[0])
		} else {
			proc.User = "unknown"
		}
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

func (proc Proc) Tags() influx.Tags {
	tags := influx.Tags{
		"name": proc.Name,
		"pid":  fmt.Sprint(proc.Pid),
		"user": proc.User,
	}

	return tags
}

func (proc Proc) Fields() influx.Fields {
	fields := influx.Fields{
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

func FindInteresting(filter Filter, debug bool) []*Proc {
	procs := make([]*Proc, 0)

	for _, pid := range AllPids() {
		proc := NewProc(pid)

		if proc == nil {
			if debug {
				fmt.Println("Failed to create Proc() for " + fmt.Sprint(pid))
			}
			continue
		}

		if len(proc.Errors) > 0 {
			if debug {
				fmt.Println("Error creating Proc() for " + fmt.Sprint(pid) + ": " + fmt.Sprint(proc.Errors[0]))
			}
			continue
		}

		if filter.Match(*proc) {
			procs = append(procs, proc)
		}
	}

	return procs
}
