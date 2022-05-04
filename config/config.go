package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigName("ServerCfg")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		switch T := err.(type) {
		case viper.ConfigFileNotFoundError:
			panic(any(fmt.Errorf("Config file not found: %w \n", T)))
		case viper.ConfigMarshalError:
			panic(any(fmt.Errorf("failing to marshal the configuration: %w \n", T)))
		case viper.ConfigParseError:
			panic(any(fmt.Errorf("config parse error: %w \n", T)))
		case viper.ConfigFileAlreadyExistsError, viper.RemoteConfigError, viper.UnsupportedConfigError, viper.UnsupportedRemoteProviderError:
			panic(any(fmt.Errorf("other error: %w \n", T)))
		default:
			fmt.Println("config read succ!")
		}
	}
}
