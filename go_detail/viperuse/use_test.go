package vipruse_test

import (
	"fmt"
	"testing"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func TestVip(t *testing.T) {
	viper.SetDefault("/home", "content")
	fmt.Println(viper.GetViper().Get("/home"))
	viper.SetConfigName("development")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../config/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println(viper.GetString("server.port"))
	viper.WatchConfig()
	go viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

}
