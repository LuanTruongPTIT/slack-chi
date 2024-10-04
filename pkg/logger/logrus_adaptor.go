package logger

import (
	"github.com/sirupsen/logrus"
)

type LogrusHandler struct {
	logger *logrus.Logger
}

func NewLogrusHandler(logger *logrus.Logger) *LogrusHandler {
	return &LogrusHandler{logger: logger}

}
