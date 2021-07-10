package app

type Market struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func NewMarket(name string) *Market {
	return &Market{Name: name}
}

func (m *Market) PrintReverseName() {}
