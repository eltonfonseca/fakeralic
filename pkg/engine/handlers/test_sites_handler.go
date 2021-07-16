package handlers

import (
	"fakeralic/pkg/engine/interfaces"
	"fakeralic/pkg/engine/models"

	"github.com/sirupsen/logrus"
)

type TestSitesHandler struct{}

func (ts *TestSitesHandler) Execute(m *models.Message) (interfaces.Handler, error) {
	logrus.Info("... Testing sites and call next handler WriteLogsHandler ...")

	return &WriteLogsHandler{}, nil
}
