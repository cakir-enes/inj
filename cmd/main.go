package main

import (
	"fmt"
	_ "gogui/injector"
	"time"

	_ "google.golang.org/grpc"
)

func fibonacci(quit <-chan int) <-chan int {
	c := make(chan int)
	x, y := 0, 1
	go func() {
		for {
			select {
			case c <- x:
				x, y = y, x+y
			case <-quit:
				fmt.Println("quit")
				return
			default:
				fmt.Println("sel2")
			}
			fmt.Println("sel")
		}
	}()
	return c
}

func main() {

	quit := make(chan int)
	_ = fibonacci(quit)
	time.Sleep(3 * time.Second)
	quit <- 1
	time.Sleep(2 * time.Second)
	// conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	// // handle err
	// defer conn.Close()
	// client := injector.NewInjectorClient(conn)
	// resp, err := client.GetAllEnumValues()

	// u, err := client.GetUser(context.Background(), &user.Empty{})
	// handle err
	// fmt.Printf("Name: %s", a)
}
