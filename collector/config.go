package collector

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"sync"
)

type Config struct {
	Compute struct {
		Freq int64 `yaml:"freq"`
	} `yaml:"compute"`
}

var config Config
var configOnce sync.Once

var logTag string

func loadConfig() {
	byteValue, err := os.ReadFile("./config/collector.yaml")
	//fmt.Println(string(byteValue))
	if err != nil {
		fmt.Println(err)
	}
	err = yaml.Unmarshal(byteValue, &config)
	//fmt.Println(config.Compute.Freq)
	if err != nil {
		fmt.Println("error parsing config.yaml")
	}
}

func init() {
	//fmt.Println("init")
	configOnce.Do(loadConfig)
}
