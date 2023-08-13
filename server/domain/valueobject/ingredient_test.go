package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIngredientCreationPositive(t *testing.T) {
	// arrange
	value := "kg"

	// act
	q, err := NewIngredient(value)

	// assert
	assert.NotNil(t, q)
	assert.Nil(t, err)
}

func TestIngredientCreationNegative(t *testing.T) {
	// arrange
	value := ""

	// act
	q, err := NewIngredient(value)

	// assert
	assert.Nil(t, q)
	assert.EqualError(t, err, "Ingredient name must not be empty")
}

func TestIngredientName(t *testing.T) {
	// arrange
	value := "potatoes"
	q, _ := NewIngredient(value)

	// act
	amount := q.Name()

	// assert
	assert.Equal(t, value, amount)
}

func TestIngredientId(t *testing.T) {
	// arrange
	value := "potatoes"
	q, _ := NewIngredient(value)

	// act
	id := q.Id()

	// assert
	assert.Equal(t, "potatoes", id)
}

func TestIngredientEqualsPositive(t *testing.T) {
	// arrange
	value := "potatoes"
	q, _ := NewIngredient(value)
	q2, _ := NewIngredient(value)

	// act
	equals := q.Equals(q2)

	// assert
	assert.True(t, equals)
}

func TestIngredientEqualsNegative(t *testing.T) {
	// arrange
	value := "potatoes"
	q, _ := NewIngredient(value)
	q2, _ := NewIngredient("tomatoes")

	// act
	equals := q.Equals(q2)

	// assert
	assert.False(t, equals)
}

func TestIngredientToString(t *testing.T) {
	// arrange
	value := "potatoes"
	q, _ := NewIngredient(value)

	// act
	s := q.ToString()

	// assert
	assert.Equal(t, "potatoes", s)
}
