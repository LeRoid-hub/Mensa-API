package models

type Day struct {
	DayName string
	Menu    []Menu
}

func (d *Day) SetDay(dayName string) {
	d.DayName = dayName
}

func (d *Day) AddMenu(menu Menu) {
	d.Menu = append(d.Menu, menu)
}
