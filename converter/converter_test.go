package main

import "testing"

func TestCelsiusToFahrenheit(t *testing.T) {
  got := celsiusToFahrenheit(0)
  want := 32.0

  if got != want {
    t.Errorf("got %.2f want %.2f", got, want)
  }
}
