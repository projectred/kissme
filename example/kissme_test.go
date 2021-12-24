package kissme

import (
	"context"
	"testing"

	"github.com/projectred/kissme"
	"github.com/projectred/kissme/example/kiss"
	"github.com/projectred/kissme/example/me"
)

type km struct {
	Honey string `kiss:"kiss"`
	Love  string `kiss:"me"`
}

func TestKissMe(t *testing.T) {
	ctx := context.WithValue(context.Background(), kiss.Kiss, "sweet")
	ctx = context.WithValue(ctx, me.Me, "forever")
	var km km
	kissme.LipPrint(ctx, &km)
	if km.Honey != "sweet" {
		t.Errorf("km.Honey should be sweety, but it is %s", km.Honey)
	}
	if km.Love != "forever" {
		t.Errorf("km.Love should be forever, but it is %s", km.Love)
	}
}

func BenchmarkKissMe(b *testing.B) {
	ctx := context.WithValue(context.Background(), kiss.Kiss, "sweet")
	ctx = context.WithValue(ctx, me.Me, "forever")
	var km km
	for i := 0; i < b.N; i++ {
		kissme.LipPrint(ctx, &km)
	}
}
