package aggregate

import (
	"testing"

	vo "domain/valueobject"

	"github.com/stretchr/testify/assert"
)

func setup() []vo.GroceryItem {
	ingredient, _ := vo.NewIngredient("apple")
	quantity, _ := vo.NewQuantity(5.0)
	unit, _ := vo.NewUnit("kg")
	item := vo.NewGroceryItem(*ingredient, *quantity, *unit)
	items := []vo.GroceryItem{*item}
	return items
}
func TestGroceryListCreationPositive(t *testing.T) {
	// arrange
	items := setup()
	// act
	list := NewGroceryList(items)

	// assert
	assert.NotNil(t, list)
}

func TestGroceryListItems(t *testing.T) {
	// arrange
	items := setup()
	item := &items[0]
	list := NewGroceryList(items)

	// act
	result := list.Items()

	// assert
	assert.NotNil(t, result)
	assert.Equal(t, 1, len(result))
	assert.Equal(t, *item, result[0])
}

func TestGroceryListAddItem(t *testing.T) {
	// arrange
	items := setup()
	list := NewGroceryList(items)

	// act
	ingredient2, _ := vo.NewIngredient("banana")
	quantity2, _ := vo.NewQuantity(2.0)
	unit2, _ := vo.NewUnit("kg")
	item2 := vo.NewGroceryItem(*ingredient2, *quantity2, *unit2)
	list.AddItem(*item2)

	// assert
	assert.NotNil(t, list)
	assert.Equal(t, 2, len(list.items))
	assert.Equal(t, *item2, list.items[1])
}

func TestGroceryListToString(t *testing.T) {
	// arrange
	items := setup()
	list := NewGroceryList(items)

	// act
	result := list.ToString()

	// assert
	assert.NotNil(t, result)
	assert.Equal(t, "5 kg apple\n", result)
}
