package config

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConfig(t *testing.T) {
	config, err := New("../config.yml")
	require.NoError(t, err)
	for _, token := range config.Tokens {
		fmt.Printf("name:%v, address:%v\n", token.Name, token.Address)
	}
	fmt.Printf("db:%+v\n", config.DB)
}
