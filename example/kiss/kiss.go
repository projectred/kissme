package kiss

import "github.com/projectred/kissme"

type kiss int

var Kiss kiss

func init() {
	kissme.Kiss("kiss", Kiss)
}
