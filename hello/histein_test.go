package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "example.com/greetings"
)

func TestHiStein(t *testing.T) {

  assert.Equal(t, "Hi, Stein. Welcome!", greetings.Hello("Stein"), "they should be equal")

}