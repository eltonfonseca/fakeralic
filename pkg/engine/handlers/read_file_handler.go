package handlers

import (
	"bufio"
	"fakeralic/pkg/engine/interfaces"
	"fakeralic/pkg/engine/models"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

type ReadFileHandler struct{}

func (rf *ReadFileHandler) Execute(m *models.Message) (interfaces.Handler, error) {
	logrus.Info("... Reading file ...")
	path, _ := filepath.Abs("hosts.txt")
	file, err := os.Open(path)

	if err != nil {
		logrus.WithError(err).Errorf("error on open file")
		return nil, err
	}

	reader := bufio.NewReader(file)

	for {
		row, err := reader.ReadString('\n')
		row = strings.TrimSpace(row)

		host := models.Host{
			Name:       row,
			Online:     false,
			StatusCode: 0,
		}

		m.Hosts = append(m.Hosts, host)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return &TestSitesHandler{}, nil
}
