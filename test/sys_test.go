package test

import (
	`testing`
	
	`github.com/chaodoing/frame/command`
)

func TestSys(t *testing.T) {
	Description, err := command.Description.Run()
	if err != nil {
		t.Error(err)
	}
	t.Log(Description)
}
