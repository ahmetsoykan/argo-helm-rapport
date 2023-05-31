package chart

import (
	"testing"
)

func TestNewHelmClient(t *testing.T) {
	_, err := NewHelmClient()
	if err != nil {
		t.Errorf("%s, test failed with error %s", "TestNewHelmClient", err)
	}
}
