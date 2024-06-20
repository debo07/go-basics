package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
 ctx := context.Background()
 // As the context has 2s timeout, it will timeout(context deadline exceeded) before slowTask is done.
 //If we change it to 3s, it will print "Operation Complete"
 ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
 // Can be done using deadline context
 // ctx, cancel := context.WithDeadline(ctx, time.Now().Add(3 * time.Second))
 defer cancel()

 go slowTask(ctx)

 //give some time for the slowTask goroutine to conclude. This is required as we are not using channel here.
 // if we don't do this, main goroutine will exit without waiting for other routines.
  time.Sleep(4 * time.Second)
}


func slowTask(ctx context.Context) {
 
 select {
 case <-time.After(3 * time.Second):
  fmt.Println("Operation complete.")
 case <-ctx.Done():
  fmt.Println("Canceled:", ctx.Err())
 }
}