package parseandgo

import (
	"errors"
	"fmt"
	"strconv"
)

type Config map[string]interface{}
type ToType int64

type Value struct {
	value interface{}
	err   error
}

func (c Config) Value(key ...string) Value {
	foundValue, err := getValueFromConfig(c, key)
	if err != nil {
		return Value{value: nil, err: err}
	}
	return Value{value: foundValue, err: nil}
}

func (v Value) Bool() (*bool, error) {
	if err := v.error(); err != nil {
		return nil, err
	} else {
		switch v.value.(type) {
		case bool:
			value := v.value.(bool)
			return &value, nil
		case string:
			value := v.value.(string)
			boolValue, parseBoolError := strconv.ParseBool(value)
			if parseBoolError != nil {
				return nil, errValueCannotBeParsedAsBool()
			}
			return &boolValue, nil
		default:
			return nil, errValueCannotBeParsedAsBool()
		}
	}
}

func (v Value) Int() (*int64, error) {
	if err := v.error(); err != nil {
		return nil, err
	} else {
		switch v.value.(type) {
		case int:
			value := v.value.(int64)
			return &value, nil
		case float64:
			value := int64(v.value.(float64))
			return &value, nil
		case string:
			value, parseIntError := strconv.ParseInt(v.value.(string), 10, 0)
			if parseIntError != nil {
				return nil, errValueCannotBeParsedAsInt()
			}
			return &value, nil
		default:
			return nil, errValueCannotBeParsedAsInt()
		}
	}
}

func (v Value) Float() (*float64, error) {
	if err := v.error(); err != nil {
		return nil, err
	} else {
		switch v.value.(type) {
		case float64:
			value := v.value.(float64)
			return &value, nil
		case string:
			value, parseFloatError := strconv.ParseFloat(v.value.(string), 0)
			if parseFloatError != nil {
				return nil, errValueCannotBeParsedAsFloat()
			}
			return &value, nil
		default:
			return nil, errValueCannotBeParsedAsFloat()
		}
	}
}

func (v Value) String() (*string, error) {
	if err := v.error(); err != nil {
		return nil, err
	} else {
		value := v.value.(string)
		return &value, nil
	}
}

func (v Value) error() error {
	if v.value == nil {
		return errInvalidValueIsNil()
	}
	return v.err
}

func getValueFromConfig(m Config, ks []string) (val interface{}, err error) {
	var ok bool
	if len(ks) == 0 {
		return nil, fmt.Errorf("config needs at least one key")
	}
	if val, ok = m[ks[0]]; !ok {
		return nil, fmt.Errorf("key not found; remaining keys: %v", ks)
	} else if len(ks) == 1 { // we've reached the final key
		return val, nil
	} else if ok = isValueConvertible(val); !ok {
		return nil, fmt.Errorf("malformed structure at %#v", val)
	} else {
		return getValueFromConfig(m, ks[1:])
	}
}

func isValueConvertible(val interface{}) bool {
	_, ok := val.(Config)
	if !ok {
		_ = val.(map[string]interface{})
		return true
	}
	return ok
}

func errKeyNotFound(key string) error {
	err := errors.New(fmt.Sprintf("given key '%s' not exists in configuration", key))
	logError(err)
	return err
}

func errInvalidValueIsNil() error {
	err := errors.New("invalid value: value is <nil>")
	logError(err)
	return err
}

func errValueCannotBeParsedAsFloat() error {
	err := errors.New("value cannot be parsed as float")
	logError(err)
	return err
}

func errValueCannotBeParsedAsInt() error {
	err := errors.New("value cannot be parsed as int")
	logError(err)
	return err
}

func errValueCannotBeParsedAsBool() error {
	err := errors.New("value cannot be parsed as bool")
	logError(err)
	return err
}
