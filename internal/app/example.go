package app

type Example struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func NewExample(name string) *Example {
	return &Example{Name: name}
}

func (e *Example) PrintReverseName() {}
