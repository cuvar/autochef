package valueobject

import (
	"errors"
	"fmt"
)

type Quantity struct {
	amount float64
}

func NewQuantity(amount float64) (*Quantity, error) {
	if amount <= 0 {
		return nil, errors.New("Amount must be positive")
	}
	return &Quantity{
		amount: amount,
	}, nil
}

func (q *Quantity) Amount() float64 {
	return q.amount
}

func (q *Quantity) Multiply(factor float64) (*Quantity, error) {
	if factor <= 0 {
		return nil, errors.New("Factor must be positive")
	}
	q, err := NewQuantity(q.amount * factor)
	if err != nil {
		return nil, err
	}
	return q, nil
}

func (q *Quantity) Equals(other *Quantity) bool {
	return q.amount == other.amount
}

func (q *Quantity) ToString() string {
	if i := int(q.amount); q.amount == float64(i) {
		s := fmt.Sprintf("%d", i)
		return s
	}
	s := fmt.Sprintf("%4.2f", q.amount)
	return s
}
