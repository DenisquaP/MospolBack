package generator_test

import (
	"mospol/auth/generator"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGeneration(t *testing.T) {
	tok, err := generator.JwtGenerator("123", "top_passw", "sec")
	require.Nil(t, err)
	if tok == "" {
		t.Fatal()
	}
}
