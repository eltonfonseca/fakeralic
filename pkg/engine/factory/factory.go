package factory

import (
	"fakeralic/pkg/engine/handlers"
	"fakeralic/pkg/engine/interfaces"
)

func NewHander() interfaces.Handler {
	return &handlers.ReadFileHandler{}
}
