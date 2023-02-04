package procs

type Filter struct {
	CPUPercent    float64
	MemoryPercent float32
	Threads       int32
	Users         []string
}

func NewFilter() Filter {
	return Filter{
		CPUPercent:    1.0,
		MemoryPercent: 1.0,
		Threads:       40,
		Users:         make([]string, 0),
	}
}

func (filter Filter) Match(p Proc) bool {
	return p.CPUPercent >= filter.CPUPercent || p.MemoryPercent >= filter.MemoryPercent || p.NumThreads >= filter.Threads
}
