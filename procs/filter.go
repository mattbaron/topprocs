package procs

// Threads and MemoryPercent are typed the way they are to make flag work :(
type Filter struct {
	CPUUsage    float64
	MemoryUsage float64
	NumThreads  int64
	MemoryVMS   uint64
	AgeMS       int64
}

func NewFilter() Filter {
	return Filter{
		CPUUsage:    1.0,
		MemoryUsage: 1.0,
		NumThreads:  50,
		MemoryVMS:   5368709120,
		AgeMS:       -1,
	}
}

func (filter Filter) Match(proc Proc) bool {
	if filter.AgeMS > 0 && proc.AgeMS < filter.AgeMS {
		return false
	}

	return proc.CPUPercent >= filter.CPUUsage ||
		proc.MemoryPercent >= float32(filter.MemoryUsage) ||
		proc.NumThreads >= int32(filter.NumThreads)
}
