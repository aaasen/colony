package colony

type Resource struct {
	Amount float64
}

func NewResource(amount float64) *Resource {
	return &Resource{
		Amount: amount,
	}
}

// Merges the second resource into the first.
// Returns both resources with their new amounts.
func AddResources(a, b *Resource) (*Resource, *Resource) {
	return NewResource(a.Amount + b.Amount), NewResource(0.0)
}

// Adds some amount to the resource.
func (self *Resource) Add(amount float64) {
	self.Amount += amount
}

// Subtracts some amount from the resource.
func (self *Resource) Subtract(amount float64) {
	self.Amount -= amount
}
