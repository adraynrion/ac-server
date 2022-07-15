package rest

type (
	Vassal struct {
		Id           uint64     `json:"id"`
		Name         string     `json:"name"`
		WorkSpeed    uint8      `json:"workSpeed"`
		InstructList []Instruct `json:"instructList,omitempty"`
	}
)

func (v *Vassal) GetName() string {
	return v.Name
}
