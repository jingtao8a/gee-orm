package log

import "testing"

func TestLog(t *testing.T) {
	SetLevel(Disabled)
	Info("Test Info")
	Error("Test Error")
}
