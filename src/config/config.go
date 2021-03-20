package config

import (
	"fmt"
	"os"

	
	"github.com/spf13/viper"
)

var (
	Conf Config
)

type Config struct {
	KafKaPort        string
	Topic            string
	Partition        int
	ApplicattionPort string
	MongoPort        string
}

func init() {

	viper.SetConfigName("conf")     // name of config file (without extension)
	viper.AddConfigPath("./config") // path to look for the config file in

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}

	err = viper.Unmarshal(&Conf)
}
