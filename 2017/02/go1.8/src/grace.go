package main

import (
	"context"
	"time"
)

func main() {
	ctx := context.Background()
	timeout := 10 * time.Second

	// START OMIT
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	server.Shutdown(ctx)
	// END OMIT
}
