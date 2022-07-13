package rest

import "math/rand"

type (
	Master struct {
		name      string
		vassals   []Vassal
		orderList []Instruct
	}
)

func (m *Master) GetName() string {
	return m.name
}

func (m *Master) AddOrder(n string, p uint64) []Instruct {
	m.orderList = append(m.orderList, Instruct{name: n, period: p})
	return m.orderList
}

func (m *Master) addVassal(n string, s uint8) []Vassal {
	m.vassals = append(m.vassals, Vassal{name: n, workSpeed: s})
	return m.vassals
}

func (m *Master) NewVassal() []Vassal {
	var n string = "GiveMeAName"
	speed := rand.Intn(11) + 1
	return m.addVassal(n, uint8(speed))
}
