// +build uint

package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCSVのプロパティが取得できること(t *testing.T) {

	sut := &Csv{}
	sut.SetHeader("csv header")
	sut.SetBody("csv body")

	assert.Equal(t, "csv header", sut.header, "")
	assert.Equal(t, "csv body", sut.body, "")
}
