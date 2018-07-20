package infrastructure

import (
	"github.com/spf13/viper"
	"path"
)

type Configuration struct {
	viper.Viper
}

// Initialize configuration file
//github.com/wenzhenxi/phalgo
func NewConfiguration(filePath string, fileName string) *Configuration {
	v := viper.New()
	v.WatchConfig()
	v.SetConfigName(fileName)
	// filePath support related path and absolute path, e.g. "/a/b" "b" "./b"
	if filePath[:1] != "/" {
		v.AddConfigPath(path.Join(GetPath(), filePath))
	} else {
		v.AddConfigPath(filePath)
	}

	// Find and Read configuration file, and handle error
	if err := c.ReadInConfig(); err != nil {
		panic(err)
	}

	return &Configuration{*v}
}
