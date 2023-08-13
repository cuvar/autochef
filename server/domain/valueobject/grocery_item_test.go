package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroceryItemCreationPositive(t *testing.T) {
	// arrange
	ingredient, _ := NewIngredient("potatoes")
	quantity, _ := NewQuantity(1)
	unit, _ := NewUnit("kg")

	// act
	q := NewGroceryItem(*ingredient, *quantity, *unit)

	// assert
	assert.NotNil(t, q)
}

func TestGroceryItemCreationNegative(t *testing.T) {
	// arrange
	ingredient, _ := NewIngredient("potatoes")
	quantity, _ := NewQuantity(1)
	unit, _ := NewUnit("kg")

	// act
	q := NewGroceryItem(*ingredient, *quantity, *unit)

	// assert
	assert.NotNil(t, q)
}

func TestGroceryItemIngredient(t *testing.T) {
	// arrange
	ingredient, _ := NewIngredient("potatoes")
	quantity, _ := NewQuantity(1)
	unit, _ := NewUnit("kg")
	groceryItem := NewGroceryItem(*ingredient, *quantity, *unit)

	// act
	returnedIngredient := groceryItem.Ingredient()
	result := returnedIngredient.Equals(ingredient)

	// assert
	assert.NotNil(t, result)
	assert.Equal(t, result, true)
}

func TestGroceryItemQuantity(t *testing.T) {
	// arrange
	ingredient, _ := NewIngredient("potatoes")
	quantity, _ := NewQuantity(1)
	unit, _ := NewUnit("kg")
	groceryItem := NewGroceryItem(*ingredient, *quantity, *unit)

	// act
	returnedQuantity := groceryItem.Quantity()
	result := returnedQuantity.Equals(quantity)

	// assert
	assert.NotNil(t, result)
	assert.Equal(t, result, true)
}

func TestGroceryItemUnit(t *testing.T) {
	// arrange
	ingredient, _ := NewIngredient("potatoes")
	quantity, _ := NewQuantity(1)
	unit, _ := NewUnit("kg")
	groceryItem := NewGroceryItem(*ingredient, *quantity, *unit)

	// act
	returnedUnit := groceryItem.Unit()
	result := returnedUnit.Equals(unit)

	// assert
	assert.NotNil(t, result)
	assert.Equal(t, result, true)
}

func TestGroceryItemEqualsPositive(t *testing.T) {
	// arrange
	ingredient, _ := NewIngredient("potatoes")
	quantity, _ := NewQuantity(1)
	unit, _ := NewUnit("kg")
	groceryItem := NewGroceryItem(*ingredient, *quantity, *unit)
	groceryItem2 := NewGroceryItem(*ingredient, *quantity, *unit)

	// act
	result := groceryItem.Equals(groceryItem2)

	// assert
	assert.Equal(t, result, true)
}

func TestGroceryItemEqualsNegativeIngredient(t *testing.T) {
	// arrange
	ingredient1, _ := NewIngredient("potatoes")
	ingredient2, _ := NewIngredient("tomatoes")
	quantity, _ := NewQuantity(1)
	unit, _ := NewUnit("kg")
	groceryItem := NewGroceryItem(*ingredient1, *quantity, *unit)
	groceryItem2 := NewGroceryItem(*ingredient2, *quantity, *unit)

	// act
	result := groceryItem.Equals(groceryItem2)

	// assert
	assert.Equal(t, result, false)
}

func TestGroceryItemEqualsNegativeQuantity(t *testing.T) {
	// arrange
	ingredient, _ := NewIngredient("potatoes")
	quantity1, _ := NewQuantity(1)
	quantity2, _ := NewQuantity(2)
	unit, _ := NewUnit("kg")
	groceryItem := NewGroceryItem(*ingredient, *quantity1, *unit)
	groceryItem2 := NewGroceryItem(*ingredient, *quantity2, *unit)

	// act
	result := groceryItem.Equals(groceryItem2)

	// assert
	assert.Equal(t, result, false)
}

func TestGroceryItemEqualsNegativeUnit(t *testing.T) {
	// arrange
	ingredient, _ := NewIngredient("potatoes")
	quantity, _ := NewQuantity(1)
	unit1, _ := NewUnit("kg")
	unit2, _ := NewUnit("l")
	groceryItem := NewGroceryItem(*ingredient, *quantity, *unit1)
	groceryItem2 := NewGroceryItem(*ingredient, *quantity, *unit2)

	// act
	result := groceryItem.Equals(groceryItem2)

	// assert
	assert.Equal(t, result, false)
}
