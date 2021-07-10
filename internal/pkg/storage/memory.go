package storage

import (
	"go-scaffold/internal/app"
)

type Memory struct{}

func (m Memory) GetType() int {
	return TypeMemory
}

func (m Memory) GetAllExamples() []app.Example {
	return []app.Example{
		{Id: 1, Name: "Shanghai"},
		{Id: 1, Name: "Beijing"},
	}
}
