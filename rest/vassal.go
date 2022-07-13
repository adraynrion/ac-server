package rest

type (
	Vassal struct {
		name         string
		workSpeed    uint8
		instructList []Instruct
	}
)

func (v *Vassal) GetName() string {
	return v.name
}
