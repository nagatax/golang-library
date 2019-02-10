package Utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test105を10で割った商が10で余りが5であること(t *testing.T) {
	expect_q := 10
	expect_r := 5

	actual_q, actual_r := Div(105, 10)

	assert.Equal(t, expect_q, actual_q, "商が異なる")
	assert.Equal(t, expect_r, actual_r, "余りが異なる")
}
