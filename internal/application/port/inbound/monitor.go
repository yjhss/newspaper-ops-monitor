package inbound

import "context"

type Monitor interface {
	RunOnce(ctx context.Context)
}
