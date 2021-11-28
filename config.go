package parseandgo

import "errors"

var ErrKeyNotFound = errors.New("given key not exists in configuration")

type Config map[string]interface{}

func (c Config) GetValueAsBool(key string) bool {
	foundKey, err := getType(key, c)
	if err != nil {
		logError(err)
	}
	return foundKey.(bool)
}

func (c Config) GetValueAsInt(key string) int {
	foundKey, err := getType(key, c)
	if err != nil {
		logError(err)
	}
	return foundKey.(int)
}

func (c Config) GetValueAsString(key string) string {
	foundKey, err := getType(key, c)
	if err != nil {
		logError(err)
	}
	return foundKey.(string)
}

func (c Config) GetValueAsFloat(key string) float64 {
	foundKey, err := getType(key, c)
	if err != nil {
		logError(err)
	}
	return foundKey.(float64)
}

func getType(_key string, config Config) (interface{}, error) {
	for key, value := range config {
		if key == _key {
			return value, nil
		}
		switch typedValue := value.(type) {
		case map[string]interface{}:
			return getType(_key, typedValue)
		}
	}
	return nil, ErrKeyNotFound
}
