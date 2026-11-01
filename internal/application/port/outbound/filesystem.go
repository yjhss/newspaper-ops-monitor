package outbound

import (
	"context"

	"github.com/yjhss/newspaper-ops-monitor/internal/application/dto"
)

type FileSystem interface {
	ScanDirectory(
		ctx context.Context,
		path string,
	) ([]dto.FileEntry, error)
}
