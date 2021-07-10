package service

import (
	"go-scaffold/internal/app"
	"go-scaffold/internal/pkg/storage"
)

type MarketService interface {
	CreateExample() app.Market
	GetMarketById(id int64) app.Market
	UpdateMarket(id int64, example app.Market) app.Market
	DeleteMarketById(id int64) app.Market
}

func NewMarketService(storage storage.Storage) MarketService {
	return &marketService{Storage: storage}
}

type marketService struct {
	Storage storage.Storage
}

func (s *marketService) CreateExample() app.Market {
	panic("implement me")
}

func (s *marketService) GetMarketById(id int64) app.Market {
	panic("implement me")
}

func (s *marketService) UpdateMarket(id int64, example app.Market) app.Market {
	panic("implement me")
}

func (s *marketService) DeleteMarketById(id int64) app.Market {
	panic("implement me")
}
