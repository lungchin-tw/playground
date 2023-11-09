package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurFuncDesc(t *testing.T) {
	t.Parallel()
	assert.NotEmpty(t, CurFuncDesc())
	t.Log(CurFuncDesc())
}

func TestCurFuncName(t *testing.T) {
	t.Parallel()
	assert.NotEmpty(t, CurFuncName())
	t.Log(CurFuncName())
}
