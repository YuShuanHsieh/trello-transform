package utilities

import (
	"os"
	"testing"
)

func testTargetPortNumber(t *testing.T) {
	if p := GetPortNumber(8181); p != 8080 {
		t.Errorf("GetPortNumber should be %d, but get %d", 8080, p)
	}
}

func testDefaultPortNumber(t *testing.T) {
	if p := GetPortNumber(8181); p != 8181 {
		t.Errorf("GetPortNumber should be %d, but get %d", 8181, p)
	}
}

func TestAllResult(t *testing.T) {
	os.Setenv("SERVER_PORT", "8080")
	testTargetPortNumber(t)
	os.Unsetenv("SERVER_PORT")
	testDefaultPortNumber(t)
}
