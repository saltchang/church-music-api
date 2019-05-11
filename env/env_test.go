package env

import (
	"testing"
)

func TestLoadEnv(t *testing.T) {
	envForTest := new(Env).loadENV()

	if envForTest.TestVar != "This env var is for testing. code:7007" {
		t.Error("Env test fails.")
	} else {
		t.Log("Env test passed!")
	}
}
