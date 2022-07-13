package rest

type (
	Person interface {
		GetName() string
	}

	Instruct struct {
		name   string
		period uint64
	}
)
