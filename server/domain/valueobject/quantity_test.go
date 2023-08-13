package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuantityCreationPositive(t *testing.T) {
	// arrange
	value := 1.0

	// act
	q, err := NewQuantity(value)
	
	// assert 
	assert.NotNil(t, q)
	assert.Nil(t, err)
}

func TestQuantityCreationNegative(t *testing.T) {
	// arrange
	value := 0.0

	// act
	q, err := NewQuantity(value)
	
	// assert 
	assert.Nil(t, q)
	assert.EqualError(t, err, "Amount must be positive")
}

func TestQuantityAmount(t *testing.T) {
	// arrange
	value := 1.0
	q, _ := NewQuantity(value)

	// act
	amount := q.Amount()
	
	// assert 
	assert.Equal(t, value, amount)
}

func TestQuantityMultiplyPositive(t *testing.T) {
	// arrange
	value := 1.0
	q, _ := NewQuantity(value)
	factor := 2.0

	// act
	q2, err := q.Multiply(factor)
	
	// assert 
	assert.NotNil(t, q2)
	assert.Nil(t, err)
	assert.Equal(t, value * factor, q2.Amount())
}

func TestQuantityMultiplyNegative(t *testing.T) {
	// arrange
	value := 1.0
	q, _ := NewQuantity(value)
	factor := 0.0

	// act
	q2, err := q.Multiply(factor)
	
	// assert 
	assert.Nil(t, q2)
	assert.EqualError(t, err, "Factor must be positive")
}

func TestQuantityEqualsPositive(t *testing.T) {
	// arrange
	value := 1.0
	q, _ := NewQuantity(value)
	q2, _ := NewQuantity(value)

	// act
	equals := q.Equals(q2)
	
	// assert 
	assert.True(t, equals)
}

func TestQuantityEqualsNegative(t *testing.T) {
	// arrange
	value := 1.0
	q, _ := NewQuantity(value)
	q2, _ := NewQuantity(value + 1)

	// act
	equals := q.Equals(q2)
	
	// assert 
	assert.False(t, equals)
}

func TestQuantityToString(t *testing.T) {
	// arrange
	value := 1.0
	q, _ := NewQuantity(value)

	// act
	s := q.ToString()
	
	// assert 
	assert.Equal(t, "1.000000", s)
}