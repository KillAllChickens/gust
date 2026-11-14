package config

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

// github.com/pelletier/go-toml/v2

type Config struct {
	Units    string  `toml:"units"`
	Lat      float64 `toml:"lat"`
	Lon      float64 `toml:"lon"`
	Location string  `toml:"location"`
}

func Default() Config {
	return Config{Units: "metric"}
}

func Path() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "gust", "config.toml"), nil
}

func Load() (Config, error) {
	cfg := Default()
	path, err := Path()
	if err != nil {
		return cfg, err
	}
	b, err := os.ReadFile(path)
	if errors.Is(err, fs.ErrNotExist) {
		return cfg, nil
	}
	if err != nil {
		return cfg, err
	}
	if err := toml.Unmarshal(b, &cfg); err != nil {
		return cfg, fmt.Errorf("parse %s: %w", path, err)
	}
	return cfg, cfg.Validate()
}

func (c Config) Validate() error {
	if c.Units != "metric" && c.Units != "imperial" {
		return fmt.Errorf("units: must be metric or imperial, got %q", c.Units)
	}
	if c.Lat < -90 || c.Lat > 90 {
		return fmt.Errorf("lat: out of range: %v", c.Lat)
	}
	if c.Lon < -180 || c.Lon > 180 {
		return fmt.Errorf("lon: out of range: %v", c.Lon)
	}
	return nil
}


func (c Config) Save() error {
    path, err := Path()
    if err != nil {
        return err
    }
    if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
        return err
    }
    b, err := toml.Marshal(c)
    if err != nil {
        return err
    }
    tmp := path + ".tmp"
    if err := os.WriteFile(tmp, b, 0o644); err != nil {
        return err
    }
    return os.Rename(tmp, path)
}
