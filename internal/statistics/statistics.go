package statistics

import "github.com/google/uuid"

type Statistics struct {
	TotalRequests uint    `json:"total_requests`
	StatusOK      bool    `jston:"status_ok"`
	Uptime        uint64  `json:"uptime"`
	StartTime     uint64  `json:"start_time"`
	Errors        []Error `json:"errors, omitempty"`
}

type Error struct {
	ID            string `json:"id"`
	Message       string `json:"message"`
	SystemMessage string `json:"system_message"`
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

func (s *Statistics) AddError(msg, sysmsg string) {
	err := Error{
		ID:            uuid.New().String(),
		Message:       msg,
		SystemMessage: sysmsg,
	}

	if len(s.Errors) >= 100 {
		s.Errors = s.Errors[1:]
	}
	s.Errors = append(s.Errors, err)
}

func (s *Statistics) GetInfo() *Statistics {
	return &Statistics{
		TotalRequests: s.TotalRequests,
		StatusOK:      s.StatusOK,
		Uptime:        s.getUptime(),
		StartTime:     s.StartTime,
	}
}
