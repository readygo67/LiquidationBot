package config

import (
	"string"
	"ioutil"
	)


type Config struct{
	RPCURL string `yaml:rpc_url`
	Network string `yaml:network`
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