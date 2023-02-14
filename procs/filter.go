package procs

// Threads and MemoryPercent are typed the way they are to make flag work :(
type Filter struct {
	CPUUsage    float64
	MemoryUsage float64
	NumThreads  int64
	MemoryVMS   uint64
}

func NewFilter() Filter {
	return Filter{
		CPUUsage:    1.0,
		MemoryUsage: 1.0,
		NumThreads:  40,
		MemoryVMS:   10000000,
	}
}

func (filter Filter) Match(p Proc) bool {
	return p.CPUPercent >= filter.CPUUsage || p.MemoryPercent >= float32(filter.MemoryUsage) || p.NumThreads >= int32(filter.NumThreads)
}
