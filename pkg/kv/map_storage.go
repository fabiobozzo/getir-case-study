package kv

import (
	"getir-case-study/pkg/utils"
	"strings"
)

type MapStorage struct {
	keys map[string]string
}

func NewMapStorage() Storage {
	return &MapStorage{
		keys: map[string]string{},
	}
}

func (m *MapStorage) Put(key, value string) error {
	// I'm enforcing as non-empty constraint, here.
	// Just a personal choice, for the assignment.
	if len(strings.TrimSpace(value)) == 0 {
		return utils.ErrEmptyValue
	}

	m.keys[key] = value

	return nil
}

func (m *MapStorage) Get(key string) (string, error) {
	value, found := m.keys[key]
	if !found {
		return "", utils.ErrNonExistingKey
	}

	return value, nil
}
