package app

import (
	"errors"
	"log/slog"
)

var ErrCannotCreateContainer = errors.New("cannot create container")

type ContainerFactoryFunc func() (*Container, error)

func NewDefaultParameters() Parameters {
	return Parameters{}
}

type Parameters struct {
	Versions map[string]string
	LogLevel slog.Level
	Debug    bool
}

type services struct {
}

func NewContainer() *Container {
	return &Container{
		Parameters: NewDefaultParameters(),
	}
}

type Container struct {
	Parameters
	services
}
