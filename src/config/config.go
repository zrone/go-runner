package config

import (
	"awesome-runner/types"
	"github.com/tal-tech/go-zero/core/conf"
	"sync"
)

var (
	Cnf  types.Config
	once sync.Once
)

func StoreConfig(configFile *string) {
	if Cnf == (types.Config{}) {
		once.Do(func() {
			conf.MustLoad(*configFile, &Cnf)
		})
	}
}
