package engine

import (
	"fakeralic/pkg/engine/factory"

	"github.com/sirupsen/logrus"
)

func Execute() error {
	var err error

	currentHandler := factory.NewHander()

	for currentHandler != nil {
		currentHandler, err = currentHandler.Execute()

		if err != nil {
			logrus.WithError(err).Errorf("chain Engine error")

			return err
		}
	}

	return nil
}
