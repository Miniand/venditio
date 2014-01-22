package config

import (
	"github.com/Miniand/venditio/core"
	"github.com/Miniand/venditio/inject"
	"os"
)

const (
	DEP_CONFIG = "config"
)

type Config struct {
	Defaults map[string]string
}

func New() *Config {
	return &Config{
		Defaults: map[string]string{},
	}
}

func Register(v *core.Venditio) {
	v.BindFactory(DEP_CONFIG, func(i inject.Injector) interface{} {
		return New()
	})
}

func (c *Config) Get(key string) string {
	val := os.Getenv(key)
	if val == "" {
		val = c.Defaults[key]
	}
	return val
}
