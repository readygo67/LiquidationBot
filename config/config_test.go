package config

import(
	"testing"
)

func TestConfig(*testing.T){
	config,_ := New(./config.yml)
	fmt.printf("%v",config)
}
