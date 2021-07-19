package handlers

import (
	"fakeralic/pkg/engine/interfaces"
	"fakeralic/pkg/engine/models"
	"net/http"

	"github.com/sirupsen/logrus"
)

type TestSitesHandler struct{}

func (ts *TestSitesHandler) Execute(m *models.Message) (interfaces.Handler, error) {
	for index, host := range m.Hosts {
		response, err := http.Get(host.Name)

		if err != nil {
			logrus.WithError(err).Error("error on make request")
			break
		}

		if response.StatusCode == 200 {
			m.Hosts[index].Online = true
		} else {
			m.Hosts[index].Online = false
		}

		m.Hosts[index].StatusCode = response.StatusCode
	}

	return &WriteLogsHandler{}, nil
}
