package internal

import "testing"

func TestParse(t *testing.T) {
	Parse("console", "PublicKey", "/Users/57block/workspace/app")
}

func TestParse1(t *testing.T) {
	SetService("PublicKey", "/Users/57block/workspace/app")
}
