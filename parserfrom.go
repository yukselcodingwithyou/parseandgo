package parseandgo

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type ParserFrom interface {
	parseFrom() []byte
}

type ParserFromURL struct {
	address string
}

func newParserFromURL(address string) ParserFromURL {
	return ParserFromURL{address: address}
}

func (parserFromURL ParserFromURL) parseFrom() []byte {
	response, err := http.Get(parserFromURL.address)
	if err != nil {
		panic(couldNotGetConfigurationFromURL(err))
	}

	body, readAllBodyError := ioutil.ReadAll(response.Body)
	defer func(Body io.ReadCloser) {
		bodyCloseError := Body.Close()
		logError(bodyCloseError)
	}(response.Body)
	logError(readAllBodyError)
	return body
}

type ParserFromFilePath struct {
	address string
}

func newParserFromFilePath(address string) ParserFromFilePath {
	return ParserFromFilePath{address: address}
}

func (parserFromFilePath ParserFromFilePath) parseFrom() []byte {
	file, fileOpenError := os.Open(parserFromFilePath.address)
	defer func(file *os.File) {
		err := file.Close()
		logError(err)
	}(file)
	logError(fileOpenError)

	body, fileReadError := ioutil.ReadAll(file)
	logError(fileReadError)
	return body
}

func newParserFrom(address string, addressType AddressType) ParserFrom {
	switch addressType {
	case URL:
		return newParserFromURL(address)
	case FILEPATH:
		return newParserFromFilePath(address)
	default:
		panic(noSuchParserDefined())
	}
}

func parseFrom(parserFrom ParserFrom) []byte {
	return parserFrom.parseFrom()
}
