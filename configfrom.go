package parseandgo

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
)

type ConfigFrom interface {
	toConfig() Config
}

func newConfigFrom(fileType FileType, body []byte) ConfigFrom {
	switch fileType {
	case JSON:
		return newConfigFromJSON(body)
	case YAML:
		return newConfigFromYAML(body)
	case ENV:
		return newConfigFromENV(body)
	case PROPERTIES:
		return newConfigFromProperties(body)
	default:
		panic(noSuchParserDefined())
	}
}

func toConfig(configFrom ConfigFrom) Config {
	return configFrom.toConfig()
}

type ConfigFromProperties struct {
	body []byte
}

func (configFromProperties ConfigFromProperties) toConfig() Config {
	return readIntoConfig(newFileReader(PROPERTIES, configFromProperties.body))
}

func newConfigFromProperties(body []byte) ConfigFromProperties {
	return ConfigFromProperties{body: body}
}

type ConfigFromENV struct {
	body []byte
}

func (configFromEnv ConfigFromENV) toConfig() Config {
	return readIntoConfig(newFileReader(ENV, configFromEnv.body))
}

func newConfigFromENV(body []byte) ConfigFromENV {
	return ConfigFromENV{body: body}
}

type ConfigFromYAML struct {
	body []byte
}

func (configFromYaml ConfigFromYAML) toConfig() Config {
	configuration := make(Config)
	err := yaml.Unmarshal(configFromYaml.body, &configuration)
	logError(err)
	return configuration
}

func newConfigFromYAML(body []byte) ConfigFromYAML {
	return ConfigFromYAML{body: body}
}

type ConfigFromJSON struct {
	body []byte
}

func (configFromJson ConfigFromJSON) toConfig() Config {
	configuration := make(Config)
	err := json.Unmarshal(configFromJson.body, &configuration)
	logError(err)
	return configuration
}

func newConfigFromJSON(body []byte) ConfigFromJSON {
	return ConfigFromJSON{body: body}
}
