package main

import "testing"

func TestEncryptMessage(t *testing.T) {
	clave := "dcj"
	mensaje := "I love prOgrAmming!"
	expected := "dcjI ldcjovdcje prdcjOgrdcjAmmdcjing!"
	got := encryptMessage(clave, mensaje)
	if got != expected {
		t.Errorf("Expected %v but got %v", expected, got)
	}
}
