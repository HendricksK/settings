package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test (t *testing.T) {
	fmt.Println("twilight is bomb")
	var test = "twilight is bomb"
	assert.Contains(t, test, "twilight")
}
