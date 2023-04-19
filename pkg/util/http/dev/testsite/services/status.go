package services

import (
	"os"
	"time"
)

type Status struct {
	Pid    int
	Uptime time.Duration
}

type StatusService struct {
	ut *UptimeService
}

func NewStatusService(
	ut *UptimeService,
) *StatusService {
	return &StatusService{
		ut: ut,
	}
}

func (svc *StatusService) Status() Status {
	return Status{
		Pid:    os.Getpid(),
		Uptime: svc.ut.Uptime(),
	}
}
