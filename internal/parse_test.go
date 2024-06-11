package internal

import "testing"

func TestParse(t *testing.T) {
	Parse("console", "Group", "/Users/57block/workspace/app")
}

func TestParse1(t *testing.T) {
	setService("PublicKey", "/Users/57block/workspace/app")
}
