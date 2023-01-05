package main

import (
	"os"
	"testing"
)

func Test_getEnv(t *testing.T) {
	want := "9999"
	key := "FOO"
	err := os.Setenv(key, want)
	if err != nil {
		t.Fatalf("failed to set env: %v", err)
	}

	if got := getEnv(key, "8080"); got != want {
		t.Errorf("getEnv() = %v, want %v", got, want)
	}
}
