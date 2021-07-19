package handlers

import (
	"fakeralic/pkg/engine/interfaces"
	"fakeralic/pkg/engine/models"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

type WriteLogsHandler struct{}

func (wl *WriteLogsHandler) Execute(m *models.Message) (interfaces.Handler, error) {
	logrus.Info("... Write logs ...")

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		logrus.WithError(err).Error("error on create or open file")
		return nil, err
	}

	for _, host := range m.Hosts {
		log := fmt.Sprintf(
			"%s - %s - Online: %s - Code: %s\n",
			time.Now().Format("02/01/2006 15:04:05"),
			host.Name,
			strconv.FormatBool(host.Online),
			strconv.FormatInt(int64(host.StatusCode), 10),
		)

		_, err := file.WriteString(log)

		if err != nil {
			logrus.WithError(err).Error("error on write file")
			break
		}
	}

	file.Close()

	return nil, nil
}
