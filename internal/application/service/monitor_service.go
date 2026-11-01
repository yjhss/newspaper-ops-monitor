package service

import (
	"context"

	"github.com/yjhss/newspaper-ops-monitor/internal/application/port/outbound"
	"github.com/yjhss/newspaper-ops-monitor/internal/domain/media"
)

type MonitorService struct {
	fileSystem outbound.FileSystem
	logger     outbound.Logger
	media      []media.Profile
}

func NewMonitorService(
	fileSystem outbound.FileSystem,
	logger outbound.Logger,
	profiles []media.Profile,
) *MonitorService {

	return &MonitorService{
		fileSystem: fileSystem,
		logger:     logger,
		media:      profiles,
	}
}

func (s *MonitorService) RunOnce(ctx context.Context) {

	for _, profile := range s.media {

		if !profile.Enabled {
			continue
		}

		for stage, path := range profile.Paths {

			files, err := s.fileSystem.ScanDirectory(ctx, path)

			if err != nil {
				s.logger.Error(
					"directory scan failed",
					"media", profile.ID,
					"stage", stage,
					"path", path,
					"error", err,
				)

				continue
			}

			fileCount := 0

			for _, file := range files {
				if !file.IsDir {
					fileCount++
				}
			}

			s.logger.Info(
				"directory scan completed",
				"media", profile.ID,
				"name", profile.Name,
				"stage", stage,
				"path", path,
				"files", fileCount,
			)
		}
	}
}
