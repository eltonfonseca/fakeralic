package interfaces

import "fakeralic/pkg/engine/models"

type Handler interface {
	Execute(m *models.Message) (Handler, error)
}
