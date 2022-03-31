package logrot

import "testing"

func TestNewLogger(t *testing.T) {
	logger := NewLogger("nick", true, true, "")
	logger.Println("123")
}
