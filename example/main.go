package main

import (
	"github.com/sheenobu/gencontext/example/enum"
	"github.com/sheenobu/gencontext/example/ptr"

	"golang.org/x/net/context"

	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = enum.NewContext(ctx, enum.Flag1)

	f := enum.Must(enum.FromContext(ctx))
	fmt.Printf("%s\n", f)

	f = enum.FromContextDefault(ctx, enum.Flag2)
	fmt.Printf("%s\n", f)

	f = enum.FromContextDefault(context.Background(), enum.Flag2)
	fmt.Printf("%s\n", f)

	ctx = ptr.NewContext(ctx, &ptr.O{
		Message: "hello",
	})

	p := ptr.Must(ptr.FromContext(ctx))

	fmt.Printf("%s\n", p.Message)

	ptr.Must(ptr.FromContext(context.Background()))

}
