// +build uint

package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test105を10で割った商が10で余りが5であること(t *testing.T) {
	expectQ := 10
	expectR := 5

	actualQ, actualR := Div(105, 10)

	assert.Equal(t, expectQ, actualQ, "商が異なる")
	assert.Equal(t, expectR, actualR, "余りが異なる")
}
