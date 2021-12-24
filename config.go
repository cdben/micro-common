package common

import (
	"fmt"
	"github.com/asim/go-micro/plugins/config/source/consul/v4"
	"go-micro.dev/v4/config"
)

// 设置配置中心
func GetConsulConfig(host string, port int64, prefix string) (config.Config, error) {
	consulSource := consul.NewSource(
		// optionally specify consul address; default to localhost:8500
		consul.WithAddress(fmt.Sprintf("%s:%d", host, port)),
		// optionally specify prefix; defaults to /micro/config
		consul.WithPrefix(prefix),
		// optionally strip the provided prefix from the keys, defaults to false
		consul.StripPrefix(true),
	)
	// config init
	conf, err := config.NewConfig()
	if err != nil {
		return conf, err
	}
	err = conf.Load(consulSource)
	return conf, err
}
