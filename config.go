package parseandgo

type Config map[string]interface{}

func (c Config) GetValueAsBool(key string) bool {
	defer func() {
		if err := recover(); err != nil {
			logPanic(err)
		}
	}()
	return c[key].(bool)
}

func (c Config) GetValueAsInt(key string) int {
	defer func() {
		if err := recover(); err != nil {
			logPanic(err)
		}
	}()
	return c[key].(int)
}

func (c Config) GetValueAsString(key string) string {
	defer func() {
		if err := recover(); err != nil {
			logPanic(err)
		}
	}()
	return c[key].(string)
}

func (c Config) GetValueAsFloat(key string) float64 {
	defer func() {
		if err := recover(); err != nil {
			logPanic(err)
		}
	}()
	return c[key].(float64)
}

func (c Config) GetValueAsConfig(key string) Config {
	defer func() {
		if err := recover(); err != nil {
			logPanic(err)
		}
	}()
	return c[key].(Config)
}
