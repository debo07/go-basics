package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go ProcessOrder(ctx)
	time.Sleep(2 * time.Second)

	// After 3 second of the ProcessOrder, main function wants to cancel it.
	fmt.Println("Main function is canceling the order processing.")
	cancel()

	time.Sleep(4 * time.Second)
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
	time.Sleep(3 * time.Second)

	if ctx.Err() != nil {
		fmt.Println("Canceled: Fetch Inventory Details")
		// fmt.Println(ctx.Err())
		return
	}
}