package gloader_test

import (
	"strings"
	"testing"

	"github.com/RecursionExcursion/gloader"
)

const envString string = `
FOO = BAR
BAZ = BIZ
`

var inMemEnv = strings.NewReader(envString)

func TestLoadEnvAndGet(t *testing.T) {
	loader := &gloader.EnvLoader{}
	err := loader.LoadEnv(inMemEnv)
	if err != nil {
		t.Fatalf("failed to load env file: %v", err)
	}

	val, err := loader.Get("FOO")
	if err != nil {
		t.Fatalf("expected FOO to be set: %v", err)
	}
	if val != "BAR" {
		t.Errorf("expected FOO to be 'bar', got '%s'", val)
	}
}

func TestMustGet_PanicsOnMissingKey(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MustGet should panic on missing key")
		}
	}()
	loader := &gloader.EnvLoader{}
	loader.LoadEnv()
	loader.MustGet("DOES_NOT_EXIST")
}

func TestGetOrFallback_ReturnsFallback(t *testing.T) {
	loader := &gloader.EnvLoader{}
	val := loader.GetOrFallback("DOES_NOT_EXIST", "fallback")
	if val != "fallback" {
		t.Errorf("expected fallback value, got '%s'", val)
	}
}

func TestGetOrDefault_StringZeroVal(t *testing.T) {
	loader := &gloader.EnvLoader{}
	val := loader.GetOrDefault("DOES_NOT_EXIST")
	if val != "" {
		t.Errorf("expected fallback value, got '%s'", val)
	}
}
