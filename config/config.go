package config

type Config struct {
	Api *ApiConfig
}

type ApiConfig struct {
	Host string
	Port int
}

func GetConfig() *Config {
	return &Config{
		Api: &ApiConfig{
			Host: "localhost",
			Port: 3000,
		},
	}
}
