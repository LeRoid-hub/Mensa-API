package models

type Mensa struct {
	Name     string
	Location string
	Days     []Day
}

func (m *Mensa) SetMensa(name string, location string) {
	m.Name = name
	m.Location = location
}

func (m *Mensa) AddDay(day Day) {
	m.Days = append(m.Days, day)
}
