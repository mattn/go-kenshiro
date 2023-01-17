package kenshiro

import (
	"context"
)

func Kenshiro(ctx context.Context) chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}
			if n == 0 {
				ch <- "あ"
			} else {
				ch <- "た"
			}
			n++
		}
	}()
	return ch
}
