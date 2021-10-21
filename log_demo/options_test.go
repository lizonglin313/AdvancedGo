package log_demo

import (
	"os"
	"testing"
)

func TestInitOptions(t *testing.T) {
	initOptions(WithLevel(DebugLevel), WithOutput(os.Stderr))
}

func TestSetOptions(t *testing.T) {
	SetOptions(WithLevel(DebugLevel), WithOutput(os.Stdout))
}