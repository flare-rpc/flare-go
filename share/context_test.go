package share

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	TheAnswer   = "Answer to the Ultimate Question of Life, the Universe, and Everything"
	MagicNumber = 42
)

func TestContext(t *testing.T) {
	flareContext := NewContext(context.Background())
	assert.NotNil(t, flareContext.Context)
	assert.NotNil(t, flareContext.tags)

	flareContext.SetValue("string", TheAnswer)
	flareContext.SetValue(42, MagicNumber)
	assert.Equal(t, MagicNumber, flareContext.Value(42))
	assert.Equal(t, TheAnswer, flareContext.Value("string"))

	flareContext.SetValue("string", TheAnswer)
	t.Log(flareContext.String())
}

func TestWithValue(t *testing.T) {
	ctx := WithValue(context.Background(), "key", "value")
	assert.NotNil(t, ctx.tags)
}

func TestWithLocalValue(t *testing.T) {
	var c Context
	c.SetValue("key", "value")

	ctx := WithLocalValue(&c, "MagicNumber", "42")

	assert.Equal(t, "value", ctx.tags["key"])
	assert.Equal(t, "42", ctx.tags["MagicNumber"])
}
