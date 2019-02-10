package Utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test整数が取得できること(t *testing.T) {
	expect := 0
	actual := 0

	sut := GenerateInteger()

	for i := 1; i <= 3; i++ {
		expect = i
		actual = sut()
		assert.Equal(t, expect, actual, "")
	}
}
