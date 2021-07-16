package handlers

import (
	"fakeralic/pkg/engine/interfaces"
	"fakeralic/pkg/engine/models"

	"github.com/sirupsen/logrus"
)

type WriteLogsHandler struct{}

func (wl *WriteLogsHandler) Execute(m *models.Message) (interfaces.Handler, error) {
	logrus.Info("... Write logs and finally ...")

	return nil, nil
}
