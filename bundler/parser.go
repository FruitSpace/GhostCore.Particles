package bundler

import (
	"encoding/json"
	"os"
	"path"
)

type GCMConfig struct {
	Script  string `json:"script"`
	Author  string `json:"author"`
	Version string `json:"version"`

	Entrypoint string `json:"entrypoint"`
	NPM        bool   `json:"npm"`
}

func NewGCMConfig(entry string) (GCMConfig, error) {
	fle, err := os.ReadFile(entry)
	if err != nil {
		return GCMConfig{}, err
	}
	conf := GCMConfig{}
	err = json.Unmarshal(fle, &conf)
	conf.Entrypoint = path.Join(path.Dir(entry), conf.Entrypoint)
	return conf, err
}

func AutoFolder(path string) (GCMConfig, error) {
	_, err := os.Stat(path)
	if err == os.ErrNotExist {
		return GCMConfig{}, err
	}
	conf, err := NewGCMConfig(path)
	if err != nil {
		return NewGCMConfig(path + "/gcm-config.json")
	}
	return conf, nil
}
