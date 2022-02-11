package config

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConfig(t *testing.T) {
	config, err := New("../config.yml")
	require.NoError(t, err)
	fmt.Printf("db:%+v\n", config.DB)
	fmt.Printf("comptroller:%+v\n", config.Comptroller)
	fmt.Printf("override:%v\n", config.Override)
	fmt.Printf("StartHeight:%v\n", config.StartHeihgt)
	fmt.Printf("PrivateKey:%v\n", config.PrivateKey)
}
