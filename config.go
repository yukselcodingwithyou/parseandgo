package parseandgo

import (
	"errors"
	"fmt"
	"reflect"
)

type Config map[string]interface{}

type Value struct {
	value interface{}
	err   error
}

func (c Config) Value(key string) Value {
	foundValue, err := getType(key, c)
	if err != nil {
		logError(err)
		return Value{value: nil, err: err}
	}
	return Value{value: foundValue, err: nil}
}

func (v Value) Bool() (*bool, error) {
	if err := v.error(reflect.String); err != nil {
		return nil, err
	} else {
		return v.value.(*bool), nil
	}
}

func (v Value) Int() (*int, error) {
	if err := v.error(reflect.String); err != nil {
		return nil, err
	} else {
		value := int(v.value.(float64))
		return &value, nil
	}
}

func (v Value) Float() (*float64, error) {
	if err := v.error(reflect.Float64); err != nil {
		return nil, err
	} else {
		return v.value.(*float64), nil
	}
}

func (v Value) String() (*string, error) {
	if err := v.error(reflect.String); err != nil {
		return nil, err
	} else {
		return v.value.(*string), nil
	}
}

func (v Value) error(t reflect.Kind) error {
	kind := reflect.ValueOf(v.value).Kind()
	if kind != t {
		return errors.New(fmt.Sprintf("type of value should be %s", kind))
	}
	return v.err
}

func getType(key string, config Config) (interface{}, error) {
	for k, v := range config {
		if key == k {
			return v, nil
		}
		switch typedValue := v.(type) {
		case map[string]interface{}:
			return getType(k, typedValue)
		}
	}
	return nil, errKeyNotFound(key)
}

func errKeyNotFound(key string) error {
	return errors.New(fmt.Sprintf("given key %s not exists in configuration", key))
}
