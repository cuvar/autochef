package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecipeStepCreationPositive(t *testing.T) {
	// arrange
	step := 1
	description := "description"

	// act
	q, err := NewRecipeStep(step, description)

	// assert
	assert.NotNil(t, q)
	assert.Nil(t, err)
}

func TestRecipeStepCreationNegativeStep(t *testing.T) {
	// arrange
	step := 0
	description := "description"

	// act
	q, err := NewRecipeStep(step, description)

	// assert
	assert.Nil(t, q)
	assert.EqualError(t, err, "step must be greater than 0")
}

func TestRecipeStepCreationNegativeDescription(t *testing.T) {
	// arrange
	step := 1
	description := ""

	// act
	q, err := NewRecipeStep(step, description)

	// assert
	assert.Nil(t, q)
	assert.EqualError(t, err, "description must not be empty")
}

func TestRecipeStepStep(t *testing.T) {
	// arrange
	step := 1
	description := "description"
	q, _ := NewRecipeStep(step, description)

	// act
	amount := q.Step()

	// assert
	assert.Equal(t, step, amount)
}

func TestRecipeStepDescription(t *testing.T) {
	// arrange
	step := 1
	description := "description"
	q, _ := NewRecipeStep(step, description)

	// act
	amount := q.Description()

	// assert
	assert.Equal(t, description, amount)
}

func TestRecipeStepEqualsPositive(t *testing.T) {
	// arrange
	step := 1
	description := "description"
	q, _ := NewRecipeStep(step, description)
	q2, _ := NewRecipeStep(step, description)

	// act
	equals := q.Equals(q2)

	// assert
	assert.True(t, equals)
}

func TestRecipeStepEqualsNegativeStep(t *testing.T) {
	// arrange
	step := 1
	description := "description"
	q, _ := NewRecipeStep(step, description)
	q2, _ := NewRecipeStep(step+1, description)

	// act
	equals := q.Equals(q2)

	// assert
	assert.False(t, equals)
}

func TestRecipeStepEqualsNegativeDescription(t *testing.T) {
	// arrange
	step := 1
	description := "description"
	description2 := "description2"
	q, _ := NewRecipeStep(step, description)
	q2, _ := NewRecipeStep(step, description2)

	// act
	equals := q.Equals(q2)

	// assert
	assert.False(t, equals)
}

func TestRecipeStepToString(t *testing.T) {
	// arrange
	step := 1
	description := "description"
	q, _ := NewRecipeStep(step, description)

	// act
	toString := q.ToString()

	// assert
	assert.Equal(t, "1. description", toString)
}
