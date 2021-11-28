package parseandgo

import (
	"github.com/yukselcodingwithyou/logandgo"
	"reflect"
)

var logger = logandgo.NewLogger(logandgo.JSON, logandgo.ERROR)

func logPanic(pnc interface{}) {
	if !isNil(pnc) {
		logger.Panic(logandgo.LogFields{
			"panic": pnc,
		})
	}
}

func logError(err error) {
	if err != nil {
		logger.Error(logandgo.LogFields{
			"error": err.Error(),
		})
	}
}

func noSuchParserDefined() string {
	panicStatement := "No such parser definition"
	logger.Panic(logandgo.LogFields{
		"panic": panicStatement,
	})
	return panicStatement
}

func couldNotGetConfigurationFromURL(err error) string {
	panicStatement := "Could not load configuration from given URL."
	logError(err)
	return panicStatement
}

func isNil(v interface{}) bool {
	return v == nil || (reflect.ValueOf(v).Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil())
}
