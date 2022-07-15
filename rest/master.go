package rest

import (
	"errors"
	"math/rand"
)

type (
	Master struct {
		Id      uint64   `json:"id"`
		Name    string   `json:"name"`
		Vassals []Vassal `json:"vassals,omitempty"`
	}
)

func (m *Master) GetName() string {
	return m.Name
}

func (m *Master) SendOrder(vId uint64, n string, p uint64) ([]Instruct, error) {
	for _, v := range m.Vassals {
		if v.Id == vId {
			v.InstructList = append(v.InstructList, Instruct{name: n, period: p})
			return v.InstructList, nil
		}
	}
	return nil, errors.New("no vassal find with given id")
}

func (m *Master) NewVassal() []Vassal {
	var n string = "GiveMeAName"
	speed := rand.Intn(11) + 1
	return m.addVassal(n, uint8(speed))
}

func (m *Master) addVassal(n string, s uint8) []Vassal {
	m.Vassals = append(m.Vassals, Vassal{Name: n, WorkSpeed: s})
	return m.Vassals
}
