package valueobject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnitCreationPositive(t *testing.T) {
	// arrange
	value := "kg"

	// act
	q, err := NewUnit(value)

	// assert
	assert.NotNil(t, q)
	assert.Nil(t, err)
}

func TestUnitCreationNegative(t *testing.T) {
	// arrange
	value := ""

	// act
	q, err := NewUnit(value)

	// assert
	assert.Nil(t, q)
	assert.EqualError(t, err, "Unit must not be empty")
}

func TestUnitValue(t *testing.T) {
	// arrange
	value := "kg"
	q, _ := NewUnit(value)

	// act
	amount := q.Value()

	// assert
	assert.Equal(t, value, amount)
}

func TestUnitId(t *testing.T) {
	// arrange
	value := "kg"
	q, _ := NewUnit(value)

	// act
	id := q.Id()

	// assert
	assert.Equal(t, "kg", id)
}

func TestUnitEqualsPositive(t *testing.T) {
	// arrange
	value := "kg"
	q, _ := NewUnit(value)
	q2, _ := NewUnit(value)

	// act
	equals := q.Equals(q2)

	// assert
	assert.True(t, equals)
}

func TestUnitEqualsNegative(t *testing.T) {
	// arrange
	value := "kg"
	q, _ := NewUnit(value)
	q2, _ := NewUnit("g")

	// act
	equals := q.Equals(q2)

	// assert
	assert.False(t, equals)
}

func TestUnitToString(t *testing.T) {
	// arrange
	value := "kg"
	q, _ := NewUnit(value)

	// act
	s := q.ToString()

	// assert
	assert.Equal(t, "kg", s)
}
