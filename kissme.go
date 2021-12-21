package kissme

import (
	"context"
	"reflect"
)

var defaultKissMe = NewKissMe("kiss")
var Kiss = defaultKissMe.Kiss
var LipPrint = defaultKissMe.LipPrint

type KissMe struct {
	tag    string
	memory map[string]interface{}
}

func NewKissMe(tag string) *KissMe {
	return &KissMe{tag: tag, memory: make(map[string]interface{})}
}

func (km KissMe) Kiss(k string, m interface{}) { km.memory[k] = m }

func (km KissMe) LipPrint(ctx context.Context, face interface{}) {
	t := reflect.TypeOf(face)
	v := reflect.ValueOf(face)
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}
	for i := 0; i < t.NumField(); i++ {
		if !v.Field(i).CanSet() {
			continue
		}
		me, ok := t.Field(i).Tag.Lookup(km.tag)
		if !ok {
			continue
		}
		honey, ok := km.memory[me]
		if !ok {
			continue
		}
		if love := ctx.Value(honey); love != nil && v.Field(i).Type().AssignableTo(reflect.TypeOf(love)) {
			v.Field(i).Set(reflect.ValueOf(love))
		}
	}
}
