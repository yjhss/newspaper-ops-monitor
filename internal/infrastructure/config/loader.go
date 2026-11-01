package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/yjhss/newspaper-ops-monitor/internal/domain/media"
	"github.com/yjhss/newspaper-ops-monitor/internal/domain/workflow"

	"gopkg.in/yaml.v3"
)

func LoadApp(path string) (*AppConfig, error) {

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config AppConfig

	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	if config.PollIntervalSeconds <= 0 {
		config.PollIntervalSeconds = 5
	}

	return &config, nil
}

func LoadMediaDirectory(path string) ([]media.Profile, error) {

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var profiles []media.Profile

	for _, entry := range entries {

		if entry.IsDir() {
			continue
		}

		ext := strings.ToLower(filepath.Ext(entry.Name()))

		if ext != ".yaml" && ext != ".yml" {
			continue
		}

		configPath := filepath.Join(path, entry.Name())

		profile, err := loadMedia(configPath)
		if err != nil {
			return nil, fmt.Errorf(
				"load media config %s: %w",
				configPath,
				err,
			)
		}

		profiles = append(profiles, profile)
	}

	return profiles, nil
}

func loadMedia(path string) (media.Profile, error) {

	data, err := os.ReadFile(path)
	if err != nil {
		return media.Profile{}, err
	}

	var config MediaConfig

	if err := yaml.Unmarshal(data, &config); err != nil {
		return media.Profile{}, err
	}

	paths := make(map[workflow.Stage]string)

	for rawStage, path := range config.Paths {

		stage, err := workflow.ParseStage(rawStage)
		if err != nil {
			return media.Profile{}, err
		}

		paths[stage] = path
	}

	return media.Profile{
		ID:      config.ID,
		Name:    config.Name,
		Enabled: config.Enabled,
		Layout:  media.LayoutType(strings.ToUpper(config.Layout.Type)),
		Ingress: media.IngressType(strings.ToUpper(config.Ingress.Type)),
		Paths:   paths,
	}, nil
}
