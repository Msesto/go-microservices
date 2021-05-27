package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Joe",
		Price: 2.69,
		SKU:   "abc-abc-abc",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
