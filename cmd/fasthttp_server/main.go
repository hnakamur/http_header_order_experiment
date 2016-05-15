package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

func requestHandler(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Add("X-Header2", "value2")
	ctx.Response.Header.Add("X-Header1", "value1")
	fmt.Fprintf(ctx, "Hello, world! Requested path is %q", ctx.Path())
}

func main() {
	var addr string
	flag.StringVar(&addr, "addr", "127.0.0.1:8080", "listen address")
	flag.Parse()

	// Create custom server.
	s := &fasthttp.Server{
		Handler: requestHandler,

		// Every response will contain 'Server: My super server' header.
		Name: "My super server",

		// Other Server settings may be set here.
	}

	// Start the server listening for incoming requests on the given address.
	//
	// ListenAndServe returns only on error, so usually it blocks forever.
	if err := s.ListenAndServe(addr); err != nil {
		log.Fatalf("error in ListenAndServe: %s", err)
	}
}
