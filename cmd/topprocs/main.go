package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mattbaron/topprocs/line"
	"github.com/mattbaron/topprocs/procs"
	"github.com/shirou/gopsutil/process"
)

var progname = filepath.Base(os.Args[0])

func Exit(code int, err any) {
	if code != 0 {
		fmt.Println(err)
	}
	os.Exit(code)
}

func GatherMetrics(measurement string, process *process.Process) *line.Line {

	// Collect basic metrics to determine if this is an interesting proc
	memPercent, err := process.MemoryPercent()
	if err != nil {
		return nil
	}

	cpuPercent, err := process.CPUPercent()
	if err != nil {
		return nil
	}

	mem, err := process.MemoryInfo()
	if err != nil {
		return nil
	}

	numThreads, err := process.NumThreads()
	if err != nil {
		return nil
	}

	if memPercent < 1 && cpuPercent < 1 && mem.Swap == 0 && numThreads < 40 {
		return nil
	}

	name, err := process.Name()
	if err != nil {
		return nil
	}

	if name == progname {
		return nil
	}

	line := line.NewLine(measurement)

	line.AddField("cpu_usage", cpuPercent)
	line.AddField("memory_usage", memPercent)
	line.AddField("memory_rss", mem.RSS)
	line.AddField("memory_vms", mem.VMS)
	line.AddField("memory_swap", mem.Swap)
	line.AddField("num_threads", numThreads)

	line.AddTag("pid", process.Pid)
	line.AddTag("name", name)

	user, err := process.Username()
	if err == nil {
		line.AddTag("user", user)
	}

	cpuTime, err := process.Times()
	if err == nil {
		line.AddField("cpu_time_iowait", cpuTime.Iowait)
		line.AddField("cpu_time_user", cpuTime.User)
		line.AddField("cpu_time_system", cpuTime.System)
	}

	return line
}

func main() {
	processes, err := procs.FindAll()
	if err != nil {
		Exit(4, err)
	}

	for _, pid := range processes {
		process, err := process.NewProcess(int32(pid))
		if err != nil {
			continue
		}

		line := GatherMetrics("topprocs", process)

		if line != nil {
			fmt.Println(line.ToString())
		}
	}
}
