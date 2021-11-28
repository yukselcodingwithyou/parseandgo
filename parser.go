package parseandgo

type FileType int64
type AddressType int64

const (
	JSON       FileType = 0
	YAML                = 1
	ENV                 = 2
	PROPERTIES          = 3
)

const (
	URL      AddressType = 0
	FILEPATH             = 1
)

type Parser interface {
	parse() Config
}

func NewParser(fileType FileType, address string, addressType AddressType) Parser {
	switch fileType {
	case JSON:
		return newJSONParser(addressType, address)
	case YAML:
		return newYAMLParser(addressType, address)
	case ENV:
		return newENVParser(addressType, address)
	case PROPERTIES:
		return newPropertiesParser(addressType, address)
	default:
		panic(noSuchParserDefined())
	}
}

func Parse(parser Parser) Config {
	return parser.parse()
}

type YAMLParser struct {
	address     string
	addressType AddressType
}

func (yp YAMLParser) parse() Config {
	return toConfig(newConfigFrom(YAML, parseFrom(newParserFrom(yp.address, yp.addressType))))
}

func newYAMLParser(addressType AddressType, address string) YAMLParser {
	return YAMLParser{address: address, addressType: addressType}
}

type JSONParser struct {
	address     string
	addressType AddressType
}

func (jp JSONParser) parse() Config {
	return toConfig(newConfigFrom(JSON, parseFrom(newParserFrom(jp.address, jp.addressType))))
}

func newJSONParser(addressType AddressType, address string) JSONParser {
	return JSONParser{address: address, addressType: addressType}
}

type ENVParser struct {
	address     string
	addressType AddressType
}

func (ep ENVParser) parse() Config {
	return toConfig(newConfigFrom(ENV, parseFrom(newParserFrom(ep.address, ep.addressType))))
}

func newENVParser(addressType AddressType, address string) ENVParser {
	return ENVParser{address: address, addressType: addressType}
}

type PropertiesParser struct {
	address     string
	addressType AddressType
}

func (pp PropertiesParser) parse() Config {
	return toConfig(newConfigFrom(PROPERTIES, parseFrom(newParserFrom(pp.address, pp.addressType))))
}

func newPropertiesParser(addressType AddressType, address string) PropertiesParser {
	return PropertiesParser{address: address, addressType: addressType}
}
