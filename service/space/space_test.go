package space

import "testing"

func TestDeleteSpace(t *testing.T) {
	err := DeleteSpace("testtt")
	if err != nil {
		t.Error(err)
	}

}

func TestNewSpace(t *testing.T) {
	err := NewSpace("testtt")
	if err != nil {
		t.Error(err)
	}
}
