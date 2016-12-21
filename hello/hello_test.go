package main

import "testing"

func TestName(t *testing.T) {
	name := getName()
	if name != "Go" {
	  t.Error ("this was not expected")
	}

}