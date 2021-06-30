package example

import (
	"go-scaffold/internal/pkg/config"
)

type Example struct {
	*config.Config
	Name string `json:"name"`
}

func NewExample(config *config.Config, name string) *Example {
	return &Example{Config: config, Name: name}
}

func (e *Example) PrintReverseName() {}
