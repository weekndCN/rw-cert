package main

import (
	"context"
	"log"
	"net/http"

	"github.com/weekndCN/rw-cert/core"
	"github.com/weekndCN/rw-cert/handler"
)

func main() {
	var ctx context.Context
	hc := core.New()
	hc.Run(ctx, "./config.yml")
	httpServer := handler.New(hc.Info(ctx))
	log.Fatal(http.ListenAndServe(":9090", httpServer.Handler()))
}
