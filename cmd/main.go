package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/Crampustallin/redisProj/internal/data_base"
	"github.com/Crampustallin/redisProj/internal/handler"
	"github.com/Crampustallin/redisProj/internal/server"
)

func main() {
	db := data_base.NewDataBase("localhost:6379")

	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()

	signalCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	h := handler.NewHandler(db)
	server := server.NewServer(h)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	fmt.Printf("server up on %s", server.Addr)

	<-signalCtx.Done()
}
