package valueobject

import (
	"errors"
	"fmt"
	"strings"
)

type Ingredient struct {
	name string
}

func NewIngredient(name string) (*Ingredient, error) {
	striped := strings.Replace(name, " ", "", -1)
	if len(striped) == 0 {
		return nil, errors.New("Ingredient name must not be empty")
	}
	return &Ingredient{
		name: striped,
	}, nil
}

func (i *Ingredient) Name() string {
	return i.name
}

func (i *Ingredient) Id() string {
	return strings.ToLower(i.name)
}

func (i *Ingredient) Equals(other *Ingredient) bool {
	return i.name == other.name
}

func (i *Ingredient) ToString() string {
	s := fmt.Sprintf("%s", i.name)
	return s
}
