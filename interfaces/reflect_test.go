package interfaces

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func takePointer(obj interface{}) {
	EnsurePointer(obj)
}

type Object struct {}

func TestEnsurePointer(t *testing.T) {
	obj := Object{}

	assert.Panics(t, func() {
		takePointer(obj)
	})

	assert.NotPanics(t, func() {
		takePointer(&obj)
	})
}

func TestEnforcePointer(t *testing.T) {
	obj := Object{}

	assert.NotPanics(t, func() {
		takePointer(EnforcePointer(obj))
	})

	assert.NotPanics(t, func() {
		takePointer(EnforcePointer(&obj))
	})
}
