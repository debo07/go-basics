package main

import (
	"context"
	"fmt"
)

// Custom type is required/better to avoid collisions between packages. Otherwise it may override others
type contextKey string

func main() {
	ctx := context.Background() // creating root empty context
	ctx = context.WithValue(ctx, contextKey("AuthzKey"), "hkhjagdja")
	makeRequest(ctx)
}

// A function which is dependent on authz value passed via Context 
func makeRequest(ctx context.Context) {
	authValue, ok := ctx.Value(contextKey("AuthzKey")).(string)
	if !ok {
		fmt.Println("Key is missing")
		return
	}
	fmt.Println(authValue)
}