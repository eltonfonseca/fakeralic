package handlers

import (
	"fakeralic/pkg/engine/interfaces"

	"github.com/sirupsen/logrus"
)

type WriteLogsHandler struct{}

func (wl *WriteLogsHandler) Execute() (interfaces.Handler, error) {
	logrus.Info("... Write logs and finally ...")

	return nil, nil
}
