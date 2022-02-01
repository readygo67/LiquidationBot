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
}
