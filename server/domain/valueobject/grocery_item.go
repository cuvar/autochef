package valueobject

type GroceryItem struct {
	ingredient Ingredient
	quantity   Quantity
	unit       Unit
}

func NewGroceryItem(ingredient Ingredient, quantity Quantity, unit Unit) *GroceryItem {
	return &GroceryItem{
		ingredient: ingredient,
		quantity:   quantity,
		unit:       unit,
	}
}

func (gi *GroceryItem) Ingredient() Ingredient {
	return gi.ingredient
}

func (gi *GroceryItem) Quantity() Quantity {
	return gi.quantity
}

func (gi *GroceryItem) Unit() Unit {
	return gi.unit
}

func (gi *GroceryItem) Equals(other *GroceryItem) bool {
	return gi.ingredient.Equals(&other.ingredient) &&
		gi.quantity.Equals(&other.quantity) &&
		gi.unit.Equals(&other.unit)
}

func (gi *GroceryItem) ToString() string {
	s := gi.quantity.ToString() + " " + gi.unit.ToString() + " " + gi.ingredient.ToString()
	return s
}
