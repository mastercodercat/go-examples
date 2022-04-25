package main

import "testing"

func TestIsMyError(t *testing.T) {
	err := ErrUserNameExist{UserName: "mastercodercat"}

	ok := IsErrUserNameExist(err)

	if !ok {
		t.Fatal("testing error")
	}
}
