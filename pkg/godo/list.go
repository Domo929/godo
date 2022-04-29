package godo

type List struct {
	Topics map[string][]Item `json:"topics,omitempty"`
}

type Item struct {
	Completion float64 `json:"completion"`
	Name       string  `json:"name"`
}

func NewList() *List {
	l := new(List)
	l.Topics = make(map[string][]Item)
	return l
}
