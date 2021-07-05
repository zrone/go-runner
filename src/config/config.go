package config

import (
	"awesome-runner/src/logr"
	"awesome-runner/types"
	"github.com/tal-tech/go-zero/core/conf"
	"gopkg.in/yaml.v2"
	"io/ioutil"
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

// parse yaml
func ParseYaml(path string, v interface{}) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		logr.Clog.Errorf("Invalid Yaml, Err: %v", err)
	} else {
		if err = yaml.Unmarshal(file, v); err != nil {
			logr.Clog.Errorf("Invalid Yaml, Err: %v", err)
		}
	}
}
