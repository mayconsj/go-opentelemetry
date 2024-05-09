package main

import (
	"context"
	"log"

	"github.com/mayconsj/go-opentelemetry/internal/otel/setup"
	"github.com/mayconsj/go-opentelemetry/internal/otel/tracing"
	"github.com/mayconsj/go-opentelemetry/internal/servicea/handler"
)

func main() {
	ctx := context.Background()
	tracer, shutdow := tracing.Start()
	defer func() {
		_ = shutdow(ctx)
	}()
	h := &handler.DefaultHandler{
		Tracer: tracer,
	}
	if err := setup.Run(h); err != nil {
		log.Fatalln(err)
	}
}
