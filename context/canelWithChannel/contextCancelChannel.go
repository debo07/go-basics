package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	done := make(chan bool)

	go func() {
		ProcessOrder(ctx)
		done <- true
	}()

	go func() {
		// Canceling after 2 seconds >> Try changing it to 3 and then 4 | If it is canceled after 4s, order processing should be done already
		time.Sleep(2 * time.Second)
		fmt.Println("Main function is canceling the order processing.")
		cancel()
	}()

	select {
	case <- done:
		fmt.Println("Order processing completed successfully")
	case <- ctx.Done():
		fmt.Println("Order processing was canceled")
	}
}

func ProcessOrder(ctx context.Context) {
	fmt.Println("Order processing...")
	// Some time consuming task here.
	time.Sleep(1 * time.Second)
	GetOrderDetails(ctx)
	if ctx.Err() != nil {
		fmt.Println("Canceled: Order processing")
		// fmt.Println(ctx.Err())
		return
	}
}

func GetOrderDetails(ctx context.Context) {
	fmt.Println("Fetching Order details")

	time.Sleep(1 * time.Second)
	GetInventoryDetails(ctx)
	if ctx.Err() != nil {
		fmt.Println("Canceled: Fetch Order Details")
		// fmt.Println(ctx.Err())
		return
	}
}

func GetInventoryDetails(ctx context.Context) {
	fmt.Println("Fetching Inventory details")

	// Some really time consuming tasks.
	time.Sleep(1 * time.Second)

	if ctx.Err() != nil {
		fmt.Println("Canceled: Fetch Inventory Details")
		// fmt.Println(ctx.Err())
		return
	}
}