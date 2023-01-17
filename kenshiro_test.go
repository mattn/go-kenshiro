package kenshiro

import (
	"context"
	"fmt"
)

func ExampleKenshiro() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := Kenshiro(ctx)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	// Output:
	// あ
	// た
	// た
	// た
	// た
	// た
	// た
	// た
	// た
	// た
}
