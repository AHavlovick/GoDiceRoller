package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestNumberofrolls(t *testing.T) {
  rolls := []int {}
  dice := 8
  rolls = GenerateRolls(6, dice)
  assert.Equal(t, dice, len(rolls) )
  }
