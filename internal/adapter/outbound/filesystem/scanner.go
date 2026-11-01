package filesystem

import (
	"context"
	"os"
	"path/filepath"

	"github.com/yjhss/newspaper-ops-monitor/internal/application/dto"
)

type Scanner struct{}

func NewScanner() *Scanner {
	return &Scanner{}
}

func (s *Scanner) ScanDirectory(
	ctx context.Context,
	path string,
) ([]dto.FileEntry, error) {

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	results := make([]dto.FileEntry, 0, len(entries))

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			continue
		}

		results = append(results, dto.FileEntry{
			Name:    entry.Name(),
			Path:    filepath.Join(path, entry.Name()),
			Size:    info.Size(),
			ModTime: info.ModTime(),
			IsDir:   entry.IsDir(),
		})
	}

	return results, nil
}
