package procs

type Filter struct {
	MinPercentCPU    float32
	MinPercentMemory float32
	MinThreads       int
	Users            []string
}

func NewFilter() *Filter {
	return &Filter{
		MinPercentCPU:    1.0,
		MinPercentMemory: 1.0,
		MinThreads:       40,
		Users:            make([]string, 0),
	}
}
