package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/shirou/gopsutil/process"
)

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

func PrintMetric(label string, value any, err any) {
	if err == nil {
		fmt.Printf("%-15s: %v\n", label, value)
	} else {
		fmt.Printf("%-15s: ERR=%v\n", label, err)
	}
}

func main() {

	processes, err := FindProcesses()
	if err != nil {
		Exit(4, err)
	} else {
		fmt.Println(cap(processes))
		fmt.Println(len(processes))
	}

	Exit(0, nil)

	if len(os.Args) < 2 {
		Exit(4, "Invalid process ID")
	}

	pid, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		Exit(3, err)
	}

	proc, err := process.NewProcess(int32(pid))
	if err != nil {
		Exit(3, err)
	}

	PrintMetric("PID", pid, nil)

	cpuPercent, err := proc.CPUPercent()
	PrintMetric("CPU Percent", cpuPercent, err)

	memPercent, err := proc.MemoryPercent()
	PrintMetric("Mem Percent", memPercent, err)

	info, err := proc.MemoryInfo()
	if err == nil {
		PrintMetric("Swap", info.Swap, nil)
		PrintMetric("RSS", info.RSS, nil)
		PrintMetric("VMS", info.VMS, nil)
	}

	cpuTime, err := proc.Times()
	if err == nil {
		PrintMetric("CPU User Time", cpuTime.User, nil)
		PrintMetric("CPU System Time", cpuTime.System, nil)
		PrintMetric("CPU Total Time", cpuTime.User+cpuTime.System, nil)
		PrintMetric("CPU IO Wait", cpuTime.Iowait, nil)
	}

	io, err := proc.IOCounters()
	if err == nil {
		PrintMetric("IO Read", io.ReadCount, err)
		PrintMetric("IO Write", io.WriteBytes, err)
	}

	numThreads, err := proc.NumThreads()
	PrintMetric("Threads", numThreads, err)

}
