package ptr

type O struct {
	Message string
}

//go:generate gencontext -type *O .
