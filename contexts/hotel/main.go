package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*5) // Set a timeout of 5 seconds for the hotel booking

	defer cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done(): // Check if the context is done or not
		fmt.Println("Hotel booking is cancelled. Timeout reached.")
		return
	case <-time.After(5 * time.Second): // Simulate a hotel booking
		fmt.Println("Hotel booking is successful.")
		return
	}
}
