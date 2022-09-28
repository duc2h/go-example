package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	go f1(ctx)
	go f2(ctx)

	<-ctx.Done()
	time.Sleep(time.Millisecond * 3)
}

func f1(ctx context.Context) {
	<-ctx.Done()

	fmt.Println("====== f1 ======")
}

func f2(ctx context.Context) {
	<-ctx.Done()

	fmt.Println("====== f2 ======")
}
