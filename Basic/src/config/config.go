package config

type ConfigDatabase struct {
	Port string ` env:"PORT" env-default:"6522"`
}

var Cfg ConfigDatabase