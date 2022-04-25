package controller

import "testing"

func TestHelloWorld(t *testing.T) {
	hello := HelloWorld("mastercodercat")
	if hello != "Hi, mastercodercat" {
		t.Errorf("Testing fail")
	}

	hello = HelloWorld("mastercodercat ")
	if hello != "Hi, mastercodercat" {
		t.Errorf("Testing fail")
	}
}
