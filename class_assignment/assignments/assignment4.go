package assignments

import (
	"fmt"
	"time"
	"math/rand/v2"
	"context"
	"sync"
)
var wg4 sync.WaitGroup

func Run4() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	wg4.Add(1)
	go func(ctx context.Context) {
		select {
		case <-time.After( time.Second * time.Duration(rand.IntN(5)) ):
			fmt.Println("work done")
			wg4.Done()
		case <-ctx.Done():
			fmt.Println("work cancel")
			wg4.Done()
		}

	}(ctx)
	wg4.Wait()
}