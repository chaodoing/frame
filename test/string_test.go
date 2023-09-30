package test

import (
	`testing`
	
	`github.com/chaodoing/frame/chars`
)

func TestString(t *testing.T) {
	t.Log(chars.Snake("developersMenu", '-'))
	t.Log(chars.Camel("create-at", '-'))
}
