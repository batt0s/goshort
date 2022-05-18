package config

import (
	"log"

	"github.com/spf13/viper"
)

var Conf *viper.Viper

// Load config files to config.Conf
func LoadConfig() error {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath(".")
	vp.AddConfigPath("..")
	vp.AddConfigPath("./config")
	vp.AddConfigPath("../config")
	err := vp.ReadInConfig()
	if err != nil {
		return err
	}
	Conf = vp
	log.Println("Loaded config from config.json")
	return nil
}
