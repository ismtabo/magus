package context_test

import (
	"testing"

	go_context "context"

	"github.com/ismtabo/magus/context"
	"github.com/stretchr/testify/assert"
)

type testKey string

const (
	testKeyFoo testKey = "foo"
)

func TestNew(t *testing.T) {
	ctx := context.New()
	assert.NotNil(t, ctx)
	assert.Equal(t, "", ctx.Cwd())
	assert.NotNil(t, ctx.Variables())
	assert.NotNil(t, ctx.Helpers())
}

func TestWith(t *testing.T) {
	expected_ctx := go_context.Background()
	expected_ctx = go_context.WithValue(expected_ctx, testKeyFoo, "bar")
	ctx := context.With(expected_ctx)
	assert.NotNil(t, ctx)
	assert.Equal(t, "bar", ctx.Value(testKeyFoo))
}

func TestWithValue(t *testing.T) {
	ctx := context.New()
	ctx = context.WithValue(ctx, testKeyFoo, "bar")
	assert.NotNil(t, ctx)
	assert.Equal(t, "bar", ctx.Value(testKeyFoo))
}

func TestContext_Cwd(t *testing.T) {
	ctx := context.New()
	assert.Equal(t, "", ctx.Cwd())
	ctx = ctx.WithCwd("/tmp")
	assert.Equal(t, "/tmp", ctx.Cwd())
}

func TestContext_Variables(t *testing.T) {
	ctx := context.New()
	assert.NotNil(t, ctx.Variables())
	ctx = ctx.WithVariables(ctx.Variables().Set("foo", "bar"))
	val, ok := ctx.Variables().Get("foo")
	assert.True(t, ok)
	assert.Equal(t, "bar", val)
}

func TestContext_Helpers(t *testing.T) {
	ctx := context.New()
	assert.NotNil(t, ctx.Helpers())
	ctx = ctx.WithHelpers(ctx.Helpers().Set("foo", "bar"))
	val, ok := ctx.Helpers().Get("foo")
	assert.True(t, ok)
	assert.Equal(t, "bar", val)
}

func TestContext_WithCwd(t *testing.T) {
	ctx := context.New()
	assert.Equal(t, "", ctx.Cwd())
	ctx = ctx.WithCwd("/tmp")
	assert.Equal(t, "/tmp", ctx.Cwd())
}

func TestContext_WithVariables(t *testing.T) {
	ctx := context.New()
	assert.NotNil(t, ctx.Variables())
	ctx = ctx.WithVariables(ctx.Variables().Set("foo", "bar"))
	val, ok := ctx.Variables().Get("foo")
	assert.True(t, ok)
	assert.Equal(t, "bar", val)
}

func TestContext_WithVariable(t *testing.T) {
	ctx := context.New()
	assert.NotNil(t, ctx.Variables())
	ctx = ctx.WithVariable("foo", "bar")
	val, ok := ctx.Variables().Get("foo")
	assert.True(t, ok)
	assert.Equal(t, "bar", val)
}

func TestContext_WithHelpers(t *testing.T) {
	ctx := context.New()
	assert.NotNil(t, ctx.Helpers())
	ctx = ctx.WithHelpers(ctx.Helpers().Set("foo", "bar"))
	val, ok := ctx.Helpers().Get("foo")
	assert.True(t, ok)
	assert.Equal(t, "bar", val)
}
