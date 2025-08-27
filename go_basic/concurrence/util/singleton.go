package util

import (
	"fmt"
	"github.com/spf13/viper"

	"sync"
)

type config struct {
	Password      string
	ServerAddress string
}

var cfg *config
var m sync.Mutex

func GetConfig() *config {
	if cfg == nil {
		m.Lock()
		defer m.Unlock()
		if cfg == nil {
			vp := viper.New()
			vp.AddConfigPath("../../config") //相对test file的路径
			vp.SetConfigName("mysql")
			vp.SetConfigType("yaml")

			fmt.Println("解析配置文件")
			if err := vp.ReadInConfig(); err != nil {
				fmt.Println(err)
				return nil
			} else {
				cfg = &config{
					Password:      vp.GetString("gift.pass"),
					ServerAddress: vp.GetString("gift.host"),
				}
			}
		}
	}
	return cfg
}
