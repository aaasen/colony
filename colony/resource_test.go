package colony

import (
	"testing"
)

func TestAddResource(t *testing.T) {
	resource := NewResource(10.0)
	input := NewResource(5.0)

	resource, input = AddResources(resource, input)

	if resource.Amount != 15.0 ||
		input.Amount != 0.0 {
		t.Fail()
	}
}
