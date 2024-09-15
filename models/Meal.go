package models

type Meal struct {
	Name       string
	Price      string
	Attributes string
}

func (m *Meal) SetMeal(name string, price string, attributes string) {
	m.Name = name
	m.Price = price
	m.Attributes = attributes
}
