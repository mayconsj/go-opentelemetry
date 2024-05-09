package main

import (
	"context"
	"log"
	"os"

	"github.com/mayconsj/go-opentelemetry/internal/otel/setup"
	"github.com/mayconsj/go-opentelemetry/internal/otel/tracing"
	"github.com/mayconsj/go-opentelemetry/internal/serviceb/handler"
	"github.com/mayconsj/go-opentelemetry/internal/serviceb/viacep"
	"github.com/mayconsj/go-opentelemetry/internal/serviceb/weather"
)

func main() {
	ctx := context.Background()
	tracer, shutdow := tracing.Start()
	defer func() {
		_ = shutdow(ctx)
	}()
	viaCepApi := &viacep.ViaCepApi{}
	weatherApi := &weather.WeatherApi{
		Key: os.Getenv("API_KEY"),
	}
	h := &handler.DefaultHandler{
		ViaCepApi:  viaCepApi,
		WeatherApi: weatherApi,
		Tracer:     tracer,
	}
	if err := setup.Run(h); err != nil {
		log.Fatalln(err)
	}
}
