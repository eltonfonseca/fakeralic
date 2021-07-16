package engine

import (
	"fakeralic/pkg/engine/factory"
	"fakeralic/pkg/engine/models"

	"github.com/sirupsen/logrus"
)

func Execute() error {
	var err error

	m := &models.Message{}
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
