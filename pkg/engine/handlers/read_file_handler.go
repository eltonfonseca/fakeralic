package handlers

import (
	"fakeralic/pkg/engine/interfaces"

	"github.com/sirupsen/logrus"
)

type ReadFileHandler struct{}

func (rf *ReadFileHandler) Execute() (interfaces.Handler, error) {
	logrus.Info("... Read file and call next handler TestSitesHandler ...")

	return &TestSitesHandler{}, nil
}
