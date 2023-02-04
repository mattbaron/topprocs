package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mattbaron/topprocs/line"
	"github.com/shirou/gopsutil/process"
)

var progname = filepath.Base(os.Args[0])

func FindProcesses() ([]int32, error) {
	procs, err := process.Pids()

	if err != nil {
		return nil, err
	}

	return procs, nil
}

func Exit(code int, err any) {
	if code != 0 {
		fmt.Println(err)
	}
	os.Exit(code)
}

func GatherMetrics(measurement string, process *process.Process) *line.Line {
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

	line := line.NewLine(measurement)

	line.AddField("cpu_usage", cpuPercent)
	line.AddField("memory_usage", memPercent)

	line.AddField("memory_rss", mem.RSS)
	line.AddField("memory_vms", mem.VMS)
	line.AddField("memory_swap", mem.Swap)

	line.AddField("num_threads", numThreads)

	line.AddTag("pid", process.Pid)

	user, err := process.Username()
	if err == nil {
		line.AddTag("user", user)
	}

	name, err := process.Name()
	if err != nil {
		return nil
	}

	if name == progname {
		return nil
	}

	line.AddTag("name", name)

	cpuTime, err := process.Times()
	if err == nil {
		line.AddField("cpu_time_iowait", cpuTime.Iowait)
		line.AddField("cpu_time_user", cpuTime.User)
		line.AddField("cpu_time_system", cpuTime.System)
	}

	return line
}

func main() {
	processes, err := FindProcesses()
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

	Exit(0, nil)

	// proc, err := process.NewProcess(int32(pid))
	// if err != nil {
	// 	Exit(3, err)
	// }

	// PrintMetric("PID", pid, nil)

	// cpuPercent, err := proc.CPUPercent()
	// PrintMetric("CPU Percent", cpuPercent, err)

	// memPercent, err := proc.MemoryPercent()
	// PrintMetric("Mem Percent", memPercent, err)

	// info, err := proc.MemoryInfo()
	// if err == nil {
	// 	PrintMetric("Swap", info.Swap, nil)
	// 	PrintMetric("RSS", info.RSS, nil)
	// 	PrintMetric("VMS", info.VMS, nil)
	// }

	// cpuTime, err := proc.Times()
	// if err == nil {
	// 	PrintMetric("CPU User Time", cpuTime.User, nil)
	// 	PrintMetric("CPU System Time", cpuTime.System, nil)
	// 	PrintMetric("CPU Total Time", cpuTime.User+cpuTime.System, nil)
	// 	PrintMetric("CPU IO Wait", cpuTime.Iowait, nil)
	// }

	// io, err := proc.IOCounters()
	// if err == nil {
	// 	PrintMetric("IO Read", io.ReadCount, err)
	// 	PrintMetric("IO Write", io.WriteBytes, err)
	// }

	// numThreads, err := proc.NumThreads()
	// PrintMetric("Threads", numThreads, err)

}
