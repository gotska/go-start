package main


import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	var variable1 string  = "helloVlad"

	p := &variable1

	*p = "helloVladhello"

	fmt.Println(&p)

	fmt.Println(variable1)
	fmt.Println(&p)



	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()

	for {
		fmt.Println("sleeping...")
		time.Sleep(10 * time.Second) // or runtime.Gosched() or similar per @misterbee
	}
}

func cleanup() {
	fmt.Println("cleanup")
}

