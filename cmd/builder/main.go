package main

import (
	"designpattern/pkg/creational/builder"
	"log"
)

func main() {
	httpBuilder := builder.NewHTTPServerBuilder().SetTLS(false)

	director := builder.NewServerDirector(
		httpBuilder,
	)

	director.ConstructSecureServer()
	httpSecureConfig := httpBuilder.GetResult()
	log.Println("[Production] HTTP config:")
	log.Printf("=> %s\n\n", httpSecureConfig.String())

	grpcBuilder := builder.NewGRPCServerBuilder()
	director.SetBuilder(grpcBuilder)
	director.ConstructStandardServer()
	grpcConfig := grpcBuilder.GetResult()
	log.Println("[Development] gRPC config:")
	log.Printf("=> %s\n\n", grpcConfig.String())
}
