package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	RPCURL        string `yaml:"rpc_url"`
	Network       string `yaml:"network"`
	Oracle        string `yaml:"oracle"`
	Comptroller   string `yaml:"comptroller"`
	PancakeRouter string `yaml:"pancake_router"`
	Liquidator    string `yaml:"liquidator"`
	PrivateKey    string `yaml:"private_key"`
	DB            string `yaml:"db"`
	StartHeihgt   uint64 `yaml:"start_heihgt"`
	Override      bool   `yaml:"override"`
}

// Setup init config
func New(path string) (*Config, error) {
	// config global config instance
	var config = new(Config)
	//h := log.StreamHandler(os.Stdout, log.TerminalFormat(true))
	//log.Root().SetHandler(h)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
