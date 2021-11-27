//go:build uint
// +build uint

package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test終了ステータスコードが取得できること(t *testing.T) {

	sut := ExitStatus(OK)
	names := sut.Names()

	assert.Equal(t, "OK", names[OK], "")
	assert.Equal(t, "ERROR", names[ERR], "")

	sut = ExitStatus(OK)
	assert.Equal(t, "OK", sut.String(), "")
	sut = ExitStatus(ERR)
	assert.Equal(t, "ERROR", sut.String(), "")
}
