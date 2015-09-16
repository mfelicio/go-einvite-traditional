package framework

import (
	"errors"
)

var factoryMap = make(map[string]func() interface{})

func GetFactory(name string) interface{} {

	if fn, ok := factoryMap[name]; ok {
		return fn()
	}

	panic(errors.New("No factory found for type " + name))
}

func SetFactory(name string, fn func() interface{}) {

	if _, ok := factoryMap[name]; !ok {
		factoryMap[name] = fn
		return
	}

	panic(errors.New("Factory already registered for type " + name))
}
