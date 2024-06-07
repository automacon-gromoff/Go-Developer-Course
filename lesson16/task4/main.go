package main

import (
	"context"
	"fmt"
)

type ctxKey string

func main() {
	ctx := context.Background()
	var someKey1 ctxKey = "some key1"
	var someKey2 ctxKey = "some key2"
	ctx = context.WithValue(ctx, someKey1, "some value1")
	ctx = context.WithValue(ctx, someKey2, "some value2")
	some(ctx)
}

func some(ctx context.Context) {
	var someKey1 ctxKey = "some key1"
	var someKey2 ctxKey = "some key2"
	fmt.Printf("%v: %v\n", someKey1, ctx.Value(someKey1))
	fmt.Printf("%v: %v", someKey2, ctx.Value(someKey2))
}
