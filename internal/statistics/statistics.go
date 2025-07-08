package statistics

type Statistics struct {
	TotalRequests uint
	StatusOK      bool
	Uptime        uint64
	StartTime     uint64
}

func New() *Statistics {
	return &Statistics{
		TotalRequests: 0,
		StatusOK:      true,
		Uptime:        0,
		StartTime:     0,
	}
}

func (s *Statistics) IncrementRequests() {
	s.TotalRequests++
}
func (s *Statistics) SetStatusOK(status bool) {
	s.StatusOK = status
}

func (s *Statistics) SetStartTime(startTime uint64) {
	if s.StartTime == 0 {
		s.StartTime = startTime
	}
}

func (s *Statistics) getUptime() uint64 {
	if s.StartTime == 0 {
		s.StartTime = uint64(0)
	}
	return uint64(0) - s.StartTime
}

func (s *Statistics) GetInfo() *Statistics {
	return &Statistics{
		TotalRequests: s.TotalRequests,
		StatusOK:      s.StatusOK,
		Uptime:        s.getUptime(),
		StartTime:     s.StartTime,
	}
}
