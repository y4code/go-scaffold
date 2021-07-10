package market

import "go-scaffold/internal/pkg/config"

type Market struct {
	*config.Config
	Name string `json:"name"`
}

func NewMarket(config *config.Config, name string) *Market {
	return &Market{Config: config, Name: name}
}

func (m *Market) PrintReverseName() {}
