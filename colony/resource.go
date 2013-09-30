package colony

import (
	"math"
)

type Resource struct {
	Amount float64
}

func NewResource(amount float64) *Resource {
	return &Resource{
		Amount: math.Max(amount, 0.0),
	}
}

// Merges the second resource into the first.
// Returns both resources with their new amounts.
func AddResources(a, b *Resource) (*Resource, *Resource) {
	return NewResource(a.Amount + b.Amount), NewResource(0.0)
}

func SubtractResources(a, b *Resource) (*Resource, *Resource) {
	return NewResource(a.Amount - b.Amount), NewResource(b.Amount)
}

// Adds some amount to the resource.
func (self *Resource) Add(amount float64) *Resource {
	return NewResource(self.Amount + amount)
}

// Subtracts some amount from the resource.
func (self *Resource) Subtract(amount float64) *Resource {
	return NewResource(self.Amount - amount)
}
