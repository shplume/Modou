package config

type ConfigReader interface {
	SetConfigPath(paths []string)
	LoadEnvConfig(file string) error
	LoadDefaultConfig(name string) error
	LoadConfig(file string) error
	GetAllConfigs() map[string]interface{}
	Get(key string) string
}

func configInit(config ConfigReader) {
	config.SetConfigPath([]string{"./config/", "."})

	if err := config.LoadEnvConfig(".env"); err != nil {
		panic(err)
	}

	if err := config.LoadDefaultConfig("config_base"); err != nil {
		panic(err)
	}

	mode := config.Get("MODE")
	if mode == "" {
		mode = "development"
	}

	configFile := "config_prod"
	if mode == "development" {
		configFile = "config_dev"
	}

	if err := config.LoadConfig(configFile); err != nil {
		panic(err)
	}
}
