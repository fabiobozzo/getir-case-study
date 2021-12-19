package kv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapStoragePut(t *testing.T) {
	testee := NewMapStorage()

	var testCases = map[string]struct {
		key     string
		value   string
		wantErr bool
	}{
		"normal key": {
			"test",
			"value",
			false,
		},
		"empty key": {
			"",
			"xx",
			true,
		},
		"empty value": {
			"yy",
			"",
			true,
		},
	}

	for _, c := range testCases {
		err := testee.Put(c.key, c.value)
		if c.wantErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestMapStorageGet(t *testing.T) {
	testee := NewMapStorage()

	testee.Put("key1", "val1")

	var testCases = map[string]struct {
		key     string
		want    string
		wantErr bool
	}{
		"existing key": {
			"key1",
			"val1",
			false,
		},
		"non-existing key": {
			"xxx",
			"",
			true,
		},
	}

	for _, c := range testCases {
		value, err := testee.Get(c.key)
		if c.wantErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, c.want, value)
		}
	}
}
