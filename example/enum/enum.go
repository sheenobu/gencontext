package enum

type Enum int

const (
	Flag1 Enum = 1
	Flag2 Enum = 2
	Flag3 Enum = 3
)

//go:generate gencontext -type Enum .
//go:generate stringer -type Enum .
