package config

var Instance *ConfigService

func init() {
	Instance = (new(ConfigService)).initConfig()
}
