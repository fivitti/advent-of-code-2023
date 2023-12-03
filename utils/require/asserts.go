package require

import (
	"reflect"
	"testing"
)

func Equal(t *testing.T, expected, actual any) {
	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func Zero(t *testing.T, actual any) {
	Equal(t, 0, actual)
}

func True(t *testing.T, actual bool) {
	Equal(t, true, actual)
}

func False(t *testing.T, actual bool) {
	Equal(t, false, actual)
}

func Len(t *testing.T, array any, expected int) {
	v := reflect.ValueOf(array)

	if v.Len() != expected {
		t.Errorf("Expected length %d, got %d", expected, v.Len())
	}
}

func Empty(t *testing.T, array any) {
	Len(t, array, 0)
}
