package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Tokens struct {
	VUSDC  string `yaml:"vUSDC"`
	VUSDT  string `yaml:"vUSDT"`
	VBUSD  string `yaml:"vBUSD"`
	VSXP   string `yaml:"vSXP"`
	VBNB   string `yaml:"vBNB"`
	VXVS   string `yaml:"vXVS"`
	VBTC   string `yaml:"vBTC"`
	VETH   string `yaml:"vETH"`
	VLTC   string `yaml:"vLTC"`
	VXRP   string `yaml:"vXRP"`
	VBCH   string `yaml:"vBCH"`
	VDOT   string `yaml:"vDOT"`
	VLINK  string `yaml:"vLINK"`
	VDAI   string `yaml:"vDAI"`
	VFIL   string `yaml:"vFIL"`
	VBETH  string `yaml:"vBETH"`
	VCAN   string `yaml:"vCAN"`
	VADA   string `yaml:"vADA"`
	VDOGE  string `yaml:"vDOGE"`
	VMATIC string `yaml:"vMATIC"`
	VCAKE  string `yaml:"vCAKE"`
	VAAVE  string `yaml:"vAAVE"`
	VTUSD  string `yaml:"vTUSD"`
	VTRX   string `yaml:"vTRX"`
}

type Config struct {
	RPCURL      string `yaml:"rpc_url"`
	Network     string `yaml:"network"`
	Unitroller  string `yaml:"unitroller"`
	Comptroller string `yaml:"comptroller"`
	Tokens      Tokens `yaml:"tokens"`
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
