package config

type AppConfig struct {
	PollIntervalSeconds int    `yaml:"poll_interval_seconds"`
	MediaConfigDir      string `yaml:"media_config_dir"`
}

type MediaConfig struct {
	ID      string `yaml:"id"`
	Name    string `yaml:"name"`
	Enabled bool   `yaml:"enabled"`

	Layout struct {
		Type string `yaml:"type"`
	} `yaml:"layout"`

	Ingress struct {
		Type string `yaml:"type"`
	} `yaml:"ingress"`

	Workflow struct {
		RIP      string `yaml:"rip"`
		Merge    string `yaml:"merge"`
		Transfer string `yaml:"transfer"`
	} `yaml:"workflow"`

	Paths map[string]string `yaml:"paths"`

	Merge struct {
		Directory     string `yaml:"directory"`
		RuntimeConfig string `yaml:"runtime_config"`
	} `yaml:"merge"`

	Reissue struct {
		Mode string `yaml:"mode"`
	} `yaml:"reissue"`
}
