package main

import "testing"

func TestVariables2(t *testing.T) {

}

func TestVariables(t *testing.T) {
  t.Run("Define a variable", func(t *testing.T) {
    var a = "initial"
    if (a != "initial") {
      t.Error("keep studying")
    }
  })

  t.Run("Define multiple variables at once", func(t *testing.T) {
    var b, c int = 1, 2
    if (b != 2 || c != 3) {
      t.Error("keep studying")
    }
  })

  t.Run("Shorthand for declaring and initializing a variable", func(t *testing.T) {
    d := 1
    if (d != 1) {
      t.Error("keep studying")
    }
  })
}

func TestVariables4(t *testing.T) {

}
