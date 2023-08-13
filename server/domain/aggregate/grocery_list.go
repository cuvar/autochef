package aggregate

import (
	vo "domain/valueobject"
)

type GroceryList struct {
	items []vo.GroceryItem
}

func NewGroceryList(items []vo.GroceryItem) *GroceryList {
	return &GroceryList{
		items: items,
	}
}

func (g *GroceryList) Items() []vo.GroceryItem {
	return g.items
}

func (g *GroceryList) AddItem(item vo.GroceryItem) {
	g.items = append(g.items, item)
}

func (g *GroceryList) ToString() string {
	name := ""
	for _, item := range g.items {
		name = name + item.ToString() + "\n"
	}
	return name
}
