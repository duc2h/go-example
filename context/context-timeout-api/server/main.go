package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	s := server{
		ctx: ctx,
	}
	r.GET("/ping", s.test2)

	fmt.Println("ok")
	r.Run(":8081")
}

type server struct {
	ctx context.Context
}

// test and test1 make timeout due to context parent call Deadline() in it's lifetime.
func (s *server) test(c *gin.Context) {
	// _, cancel := context.WithTimeout(s.ctx, time.Second*2)
	// defer cancel()
	fmt.Println("okkkkkkkkkkk")
	slow(s.ctx)

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func (s *server) test1(c *gin.Context) {
	ctx, cancel := context.WithTimeout(s.ctx, time.Second*20)
	defer cancel()
	fmt.Println("okkkkkkkkkkk")
	slow(ctx)

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func (s *server) test2(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, time.Second*20)
	defer cancel()
	fmt.Println("okkkkkkkkkkk")
	slow(ctx)

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func slow(ctx context.Context) {
	// time.Sleep(time.Second * 3)
	go timeout()
	<-ctx.Done()
	fmt.Println("slowwwwwwwww")
}

func timeout() {
	time.Sleep(time.Second * 10)
	fmt.Println("timeouttttttt")
}
