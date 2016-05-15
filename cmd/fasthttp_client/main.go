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
	flag.StringVar(&addr, "addr", "127.0.0.1:8080", "host address")
	var targetURLStr string
	flag.StringVar(&targetURLStr, "target-url", "http://example.com/world", "target URL")
	flag.Parse()

	c := &fasthttp.HostClient{
		Addr: addr,
	}
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(targetURLStr)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	err := c.Do(req, resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("header:\n%s", string(resp.Header.String()))
	fmt.Printf("body:\n%s", string(resp.Body()))
}
