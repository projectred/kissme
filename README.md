# KISSME

KISSME is a Go package that supports a easy way to manager the tag and inject the struct with context.

By the way, KISSME was wrote in a hard day, a man who fell in love, and could not help himself missing the other. Can you kiss me more!!!

Getting Started
===============

## Installing

To start using KISSME, install Go and run `go get`:

```sh
$ go get -u github.com/projectred/kissme
```

## Regist

```golang
type me int

var TagMe me

func init() {
	kissme.Kiss("me", TagMe)
}
```

## Run

```golang
type km struct {
	Kiss  string `kiss:"me"`
}

func TestKissMe(t *testing.T) {
	ctx := context.WithValue(context.Background(), TagMe, "more")
	var km km
	kissme.LipPrint(ctx, &km)
	if km.Kiss != "more" {
		t.Errorf("km.Kiss should be 'more', but it is %s", km.Kiss)
	}
}
``` 