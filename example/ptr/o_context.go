// generated by gencontext -type *O .; DO NOT EDIT

package ptr

import (
	"golang.org/x/net/context"
)

type keyType int

var key keyType

// NewContext creates a new context with the given *O
func NewContext(ctx context.Context, val *O) context.Context {
	return context.WithValue(ctx, key, val)
}

// FromContext loads the *O from the context
func FromContext(ctx context.Context) (val *O, ok bool) {
	val, ok = ctx.Value(key).(*O)
	return
}

// Must returns the given type or panics if the ok boolean is false
func Must(val *O, ok bool) *O {
	if !ok {
		panic("Can not get *O from context")
	}

	return val
}

// FromContextDefault returns the given FromContext results or, if empty, returns the default value
func FromContextDefault(ctx context.Context, def *O) *O {
	val, ok := ctx.Value(key).(*O)
	if !ok {
		return def
	}

	return val
}