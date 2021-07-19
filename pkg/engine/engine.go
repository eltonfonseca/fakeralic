package engine

import (
	"fakeralic/pkg/engine/factory"
	"fakeralic/pkg/engine/models"

	"github.com/sirupsen/logrus"
)

func Execute(m *models.Message) error {
	var err error
	currentHandler := factory.NewHander()

	for currentHandler != nil {
		currentHandler, err = currentHandler.Execute(m)

		if err != nil {
			logrus.WithError(err).Errorf("chain Engine error")

			return err
		}
	}

	return nil
}
