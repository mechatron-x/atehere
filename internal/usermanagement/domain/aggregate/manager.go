package aggregate

import (
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/usermanagement/domain/valueobject"
)

type Manager struct {
	core.Aggregate
	email       valueobject.Email
	password    valueobject.Password
	fullName    valueobject.FullName
	phoneNumber valueobject.PhoneNumber
}

func NewManager() *Manager {
	return &Manager{
		Aggregate: core.NewAggregate(),
	}
}

func (m *Manager) Email() valueobject.Email {
	return m.email
}

func (m *Manager) Password() valueobject.Password {
	return m.password
}

func (m *Manager) FullName() valueobject.FullName {
	return m.fullName
}

func (m *Manager) PhoneNumber() valueobject.PhoneNumber {
	return m.phoneNumber
}

func (m *Manager) SetEmail(email valueobject.Email) {
	m.email = email
}

func (m *Manager) SetPassword(password valueobject.Password) {
	m.password = password
}

func (m *Manager) SetFullName(fullName valueobject.FullName) {
	m.fullName = fullName
}

func (m *Manager) SetPhoneNumber(phoneNumber valueobject.PhoneNumber) {
	m.phoneNumber = phoneNumber
}
