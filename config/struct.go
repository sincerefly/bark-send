package config

type Config struct {
	Service struct {
		BaseUrl string `mapstructure:"base_url"`
	}
	Users []string
}
