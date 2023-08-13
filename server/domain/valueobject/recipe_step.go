package valueobject

import (
	"errors"
	"fmt"
	"strings"
)

type RecipeStep struct {
	step        int
	description string
}

func NewRecipeStep(step int, description string) (*RecipeStep, error) {
	if step <= 0 {
		return nil, errors.New("step must be greater than 0")
	}

	var striped = strings.TrimSpace(description)
	if len(striped) == 0 {
		return nil, errors.New("description must not be empty")
	}

	return &RecipeStep{
		step:        step,
		description: striped,
	}, nil
}

func (rs *RecipeStep) Step() int {
	return rs.step
}

func (rs *RecipeStep) Description() string {
	return rs.description
}

func (rs *RecipeStep) Equals(other *RecipeStep) bool {
	return rs.step == other.step &&
		rs.description == other.description
}

func (rs *RecipeStep) ToString() string {
	return fmt.Sprintf("%d. %s", rs.step, rs.description)
}
