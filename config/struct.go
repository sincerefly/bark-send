package config

type Config struct {
	Service struct {
		BaseUrl string `mapstructure:"base_url"`
	}
	Log Log
	Users []string
}

type Log struct {
	Path string
}