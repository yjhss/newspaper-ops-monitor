package scheduler

import (
	"context"
	"time"

	"github.com/yjhss/newspaper-ops-monitor/internal/application/port/inbound"
)

type PollingScheduler struct {
	monitor  inbound.Monitor
	interval time.Duration
}

func NewPollingScheduler(
	monitor inbound.Monitor,
	interval time.Duration,
) *PollingScheduler {

	return &PollingScheduler{
		monitor:  monitor,
		interval: interval,
	}
}

func (s *PollingScheduler) Run(ctx context.Context) {

	s.monitor.RunOnce(ctx)

	ticker := time.NewTicker(s.interval)
	defer ticker.Stop()

	for {
		select {

		case <-ctx.Done():
			return

		case <-ticker.C:
			s.monitor.RunOnce(ctx)
		}
	}
}
