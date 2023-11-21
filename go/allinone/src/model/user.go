package model

import (
	"fmt"
)

type EnumGender int

func (g EnumGender) CheckValid() error {
	if (g <= EG_MIN) || (g >= EG_MAX) {
		return fmt.Errorf("Value %v is out of EnumGender", g)
	}

	return nil
}

const (
	EG_MIN EnumGender = iota
	EG_MALE
	EG_FEMALE
	EG_MAX
)

type User struct {
	name     string
	height   int
	gender   EnumGender
	numDates int
}

func (u *User) CheckValid() error {
	if u == nil {
		return fmt.Errorf("User is nil")
	} else if len(u.Name()) == 0 {
		return fmt.Errorf("User's Name is nil")
	} else if u.Gender() == 0 {
		return fmt.Errorf("User's Gender is nil")
	}

	return nil
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Height() int {
	return u.height
}

func (u *User) Gender() EnumGender {
	return u.gender
}

func (u *User) NumDates() int {
	return u.numDates
}

func (u *User) DecreaseNumDates() int {
	u.numDates--
	return u.NumDates()
}

func (u *User) Clone() *User {
	copy := *u
	return &copy
}

func NewUser(
	name string,
	ht int,
	gender EnumGender,
	numDates int,
) *User {
	return &User{
		name:     name,
		height:   ht,
		gender:   gender,
		numDates: numDates,
	}
}
