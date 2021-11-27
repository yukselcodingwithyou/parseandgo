package parseandgo

import "github.com/yukselcodingwithyou/logandgo"

var logger = logandgo.NewLogger(logandgo.JSON, logandgo.ERROR)

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
