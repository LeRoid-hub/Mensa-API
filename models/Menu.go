package models

type Menu struct {
	Name string
	Meal []Meal
}

func (m *Menu) SetMenu(name string) {
	m.Name = name
}

func (m *Menu) AddMeal(meal Meal) {
	m.Meal = append(m.Meal, meal)
}
