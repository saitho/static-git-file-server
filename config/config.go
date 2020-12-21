package config

import (
	"os"

	"github.com/creasty/defaults"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Git struct {
		Url       string
		WorkDir   string `yaml:"work_dir"`
		CacheTime int    `yaml:"cache_time"`
	}
	Display struct {
		Index struct {
			ShowBranches bool   `default:"true" yaml:"show_branches"`
			ShowTags     bool   `default:"true" yaml:"show_tags"`
			TagsOrder    string `default:"desc" yaml:"tags_order"`
		}
	}
	Files []string
}

func LoadConfig(configPath string) (*Config, error) {
	cfg := &Config{}
	defaults.Set(cfg)

	f, err := os.Open(configPath)
	if err != nil {
		return cfg, err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}
