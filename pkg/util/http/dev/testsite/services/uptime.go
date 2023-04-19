package services

import (
	"time"
)

type UptimeService struct {
	st time.Time
}

func NewUptimeService() *UptimeService {
	return &UptimeService{
		st: time.Now(),
	}
}

func (svc UptimeService) Uptime() time.Duration {
	return time.Now().Sub(svc.st)
}
