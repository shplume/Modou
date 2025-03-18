package config

type ConfigReader interface {
	SetConfigPath(paths []string)
	LoadEnvConfig(file string) error
	LoadDefaultConfig(name string) error
	LoadConfig(file string) error
	GetAllConfigs() map[string]interface{}
	Get(key string) string
}
