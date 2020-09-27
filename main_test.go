package main

import (
	"testing"
)

func TestTendoMain(t *testing.T) {
	defer handlePanic(t, "Tendo")

	main()
}

func handlePanic(t *testing.T, structName string) {
	recovery := recover()
	if recovery != nil {
		t.Errorf("%s function should not panic.", structName)
	}
}
