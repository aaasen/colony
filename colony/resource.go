package colony

type Resource struct {
	Amount float64
}

func NewResource(amount float64) *Resource {
	return &Resource{
		Amount: amount,
	}
}
