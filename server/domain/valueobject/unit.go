package valueobject

import (
	"errors"
	"fmt"
	"strings"
)

type Unit struct {
	value string
}

func NewUnit(value string) (*Unit, error) {
	striped := strings.TrimSpace(value)
	if len(striped) == 0 {
		return nil, errors.New("Unit must not be empty")
	}
	return &Unit{
		value: striped,
	}, nil
}

func (u *Unit) Value() string {
	return u.value
}

func (u *Unit) Id() string {
	return strings.ToLower(u.value)
}

func (u *Unit) Equals(other *Unit) bool {
	return u.value == other.value
}

func (u *Unit) ToString() string {
	s := fmt.Sprintf("%s", u.value)
	return s
}
