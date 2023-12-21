package model

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
)

type Match struct {
	users    map[string]*User
	muRemove sync.Mutex
}

var defaultMatchService *Match

func init() {
	defaultMatchService = NewMatch()
}

func DefaultMatchService() *Match {
	return defaultMatchService
}

func EmptyDefaultMatchService() {
	defaultMatchService = NewMatch()
}

func NewMatch() *Match {
	return &Match{
		users: make(map[string]*User),
	}
}

func (m *Match) AddUser(user *User) error {
	if err := user.CheckValid(); err != nil {
		return err
	} else if m.getUserByName(user.Name()) != nil {
		return fmt.Errorf("User %v already exists.", user.Name())
	}

	m.users[user.name] = user.Clone()
	return nil
}

func (m *Match) FindPossiblePeople(source *User) []string {
	result := []string{}
	for _, target := range m.users {
		if m.checkMatchCondition(source, target) == true {
			result = append(result, target.Name())
		}
	}

	return result
}

func (m *Match) FindPossiblePeopleByName(user string) []string {
	result := []string{}

	source := m.getUserByName(user)
	if source == nil {
		return result
	} else {
		return m.FindPossiblePeople(source)
	}
}

func (m *Match) RandomMatch(source *User) *User {
	result := m.FindPossiblePeople(source)
	if len(result) == 0 {
		return nil
	}

	return m.getUserByName(result[rand.Intn(len(result))]).Clone()
}

func (m *Match) DecreaseNumDatesAndRemove(name string) error {
	if v, ok := m.users[name]; ok {
		if v.DecreaseNumDates() <= 0 {
			m.DeleteUser(name)
		}
	}

	return nil
}

func (m *Match) DeleteUser(name string) {
	m.muRemove.Lock()
	defer m.muRemove.Unlock()

	delete(m.users, name)
}

func (m *Match) checkMatchCondition(source, target *User) bool {
	if source.CheckValid() != nil {
		return false
	} else if target.CheckValid() != nil {
		return false
	} else if (source.NumDates() <= 0) || (target.NumDates() <= 0) {
		return false
	} else if strings.Compare(source.Name(), target.Name()) == 0 {
		return false
	} else if source.Gender() == target.Gender() {
		return false
	} else if (source.Gender() == EG_MALE) && (source.Height() <= target.Height()) {
		return false
	} else if (source.Gender() == EG_FEMALE) && (source.Height() >= target.Height()) {
		return false
	}

	return true
}

func (m *Match) getUserByName(name string) *User {
	if v, ok := m.users[name]; ok {
		return v
	}

	return nil
}
