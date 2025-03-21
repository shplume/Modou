package config

import (
	"sync"

	"github.com/spf13/viper"
)

type DefaultConfigReader struct {
	viper *viper.Viper
}

var (
	defaultConfigReaderInstance *DefaultConfigReader
	once                        sync.Once
)

func GetDefaultConfigReader() *DefaultConfigReader {
	once.Do(func() {
		defaultConfigReaderInstance = &DefaultConfigReader{
			viper: viper.New(),
		}

		configInit(defaultConfigReaderInstance)
	})
	return defaultConfigReaderInstance
}

func (v *DefaultConfigReader) LoadEnvConfig(file string) error {
	v.viper.SetConfigFile(file)

	if err := v.viper.ReadInConfig(); err != nil {
		return err
	}

	for key, val := range v.viper.AllSettings() {
		v.viper.Set(key, val)
	}

	return nil
}

func (v *DefaultConfigReader) LoadDefaultConfig(name string) error {
	v.viper.SetConfigName(name)
	if err := v.viper.ReadInConfig(); err != nil {
		return err
	}

	for key, val := range v.viper.AllSettings() {
		v.viper.SetDefault(key, val)
	}

	return nil
}

func (v *DefaultConfigReader) SetConfigPath(paths []string) {
	for _, path := range paths {
		v.viper.AddConfigPath(path)
	}
}

func (v *DefaultConfigReader) LoadConfig(name string) error {
	v.viper.SetConfigName(name)

	if err := v.viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func (v *DefaultConfigReader) Get(key string) string {
	return v.viper.GetString(key)
}

func (v *DefaultConfigReader) GetAllConfigs() map[string]interface{} {
	return v.viper.AllSettings()
}
