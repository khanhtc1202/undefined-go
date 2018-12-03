package main

import "reflect"

type TypeMapper interface {
	Map(interface{}) TypeMapper
}

type injector struct {
	values map[reflect.Type]reflect.Value
}

func (i *injector) Map(val interface{}) TypeMapper {
	i.values[reflect.TypeOf(val)] = reflect.ValueOf(val)
	return i
}
