package handlers

import (
	"fakeralic/pkg/engine/interfaces"

	"github.com/sirupsen/logrus"
)

type TestSitesHandler struct{}

func (ts *TestSitesHandler) Execute() (interfaces.Handler, error) {
	logrus.Info("... Testing sites and call next handler WriteLogsHandler ...")

	return &WriteLogsHandler{}, nil
}
