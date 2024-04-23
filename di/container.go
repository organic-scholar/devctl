package di

import (
	"fmt"
	"reflect"
)

type Container struct {
	defs map[reflect.Type]any
}

func (c *Container) Put(val any) {
	valType := reflect.TypeOf(val)
	c.defs[valType] = val
}

func Get[T any](c *Container) T {
	t := reflect.TypeFor[T]()
	if v, ok := c.defs[t]; ok {
		return v.(T)
	}
	panic(fmt.Errorf("unable to resolve type %s", t.Name()))
}

func NewContainer() *Container {
	return &Container{
		defs: make(map[reflect.Type]any),
	}
}

func (c *Container) Provide(fn interface{}) {
	fnValue := reflect.ValueOf(fn)
	fnType := fnValue.Type()
	if fnType.Kind() != reflect.Func {
		panic(fmt.Errorf("expected a function, got %s", fnType.String()))
	}
	args := make([]reflect.Value, fnType.NumIn())
	for i := 0; i < fnType.NumIn(); i++ {
		paramType := fnType.In(i)
		if val, ok := c.defs[paramType]; ok {
			args[i] = reflect.ValueOf(val)
		} else {
			panic(fmt.Errorf("unable to resolve type %s", paramType.Name()))
		}
	}
	for _, returnValue := range fnValue.Call(args) {
		c.defs[returnValue.Type()] = returnValue.Interface()
	}
}
