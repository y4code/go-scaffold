package storage

import (
	"github.com/y4code/go-scaffold/internal/app"
)

const (
	TypeMemory = iota
	TypeFile
	TypeDb
)

type Storage interface {
	GetType() int
	GetAllExamples() []app.Example
}

func NewStorage(storageType int) Storage {
	switch storageType {
	case TypeMemory:
		return Memory{}
	case TypeFile:
	//return nil
	case TypeDb:
		//return nil
	}
	return nil
}
