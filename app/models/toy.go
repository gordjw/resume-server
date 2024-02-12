package models

type Toy struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	ToyType string `json:"type"`
	Data    string `json:"data"`
}

func NewToy(id, name, toyType, data string) Toy {
	return Toy{
		Id:      id,
		Name:    name,
		ToyType: toyType,
		Data:    data,
	}
}
