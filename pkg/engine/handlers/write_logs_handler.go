package handlers

import (
	"fakeralic/pkg/engine/interfaces"
	"fakeralic/pkg/engine/models"
	"fmt"

	"github.com/sirupsen/logrus"
)

type WriteLogsHandler struct{}

func (wl *WriteLogsHandler) Execute(m *models.Message) (interfaces.Handler, error) {
	logrus.Info("... Write logs and finally ...")

	for _, host := range m.Hosts {
		fmt.Printf("%+v\n", host)
	}

	return nil, nil
}
