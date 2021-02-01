package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/robfig/cron"
	"github.com/weekndCN/rw-cert/core"
	"github.com/weekndCN/rw-cert/handler"
)

func main() {
	var ctx context.Context
	hc := core.New()
	hc.Run(ctx, "./config.yml")
	httpServer := handler.New(hc.Info(ctx))
	// cron check cert daily
	c := cron.New()
	c.AddFunc("@midnight", func() {
		var ctx context.Context
		hc.Check(ctx)
		timeNow := time.Now()
		fmt.Printf("Certs info updated: %s\n", timeNow)
	})

	c.Start()
	defer c.Stop()

	log.Fatal(http.ListenAndServe(":9090", httpServer.Handler()))
}
