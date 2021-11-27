package parseandgo

import (
	"bufio"
	"bytes"
	"strings"
	"unicode"
)

type FileReader interface {
	readIntoConfig() Config
}

func newFileReader(fileType FileType, body []byte) FileReader {
	switch fileType {
	case ENV:
		return EnvFileReader{body: body}
	case PROPERTIES:
		return PropertiesFileReader{body: body}
	default:
		panic(noSuchParserDefined())
	}
}

func readIntoConfig(reader FileReader) Config {
	return reader.readIntoConfig()
}

type EnvFileReader struct {
	body []byte
}

func (envFileReader EnvFileReader) readIntoConfig() Config {
	return readFileLineByLine(getFile(envFileReader.body), "#")
}

type PropertiesFileReader struct {
	body []byte
}

func (propertiesFileReader PropertiesFileReader) readIntoConfig() Config {
	return readFileLineByLine(getFile(propertiesFileReader.body), "#")
}

func getFile(body []byte) *bufio.Scanner {
	reader := bytes.NewReader(body)
	return bufio.NewScanner(reader)
}

func readFileLineByLine(scanner *bufio.Scanner, commentToken string) Config {
	configuration := make(Config)
	for scanner.Scan() {
		line := scanner.Text()
		configuration = processLine(line, configuration, commentToken)
	}
	return configuration
}

func processLine(line string, config Config, commentToken string) Config {
	if isLineNotCommented(line, commentToken) {
		spacesStripped := stripSpacesFromLine(line)
		key, value := splitLine(spacesStripped)
		config[key] = value
	}
	return config
}

func isLineNotCommented(line string, substr string) bool {
	return !strings.Contains(line, substr)
}

func stripSpacesFromLine(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}

func splitLine(line string) (string, string) {
	keyAndValue := strings.Split(line, "=")
	return keyAndValue[0], keyAndValue[1]
}
