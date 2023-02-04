package procs

// Threads and MemoryPercent are typed the way they are to make flag work :(
type Filter struct {
	CPUPercent    float64
	MemoryPercent float64
	Threads       int64
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
	return p.CPUPercent >= filter.CPUPercent || p.MemoryPercent >= float32(filter.MemoryPercent) || p.NumThreads >= int32(filter.Threads)
}
