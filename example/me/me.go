package me

import "github.com/projectred/kissme"

type me int

var Me me

func init() {
	kissme.Kiss("me", Me)
}
