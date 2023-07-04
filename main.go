package main

import (
	"context"

	ga "google.golang.org/api/analyticsdata/v1beta"
)

func main() {
	ga.NewService(context.Background())
}
