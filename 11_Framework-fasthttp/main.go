package main

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

var gae *gaeInfo

func indexHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/":
		fmt.Fprint(ctx, "Working Golang App")
	default:
		ctx.Error("404 Not Found", fasthttp.StatusNotFound)
	}
}

func main() {
	// Load the Environment Info from Appengine
	gae = getAppengineEnv()

	log.Printf("Listening on Port %s : Server", gae.PORT)
	log.Fatal(fasthttp.ListenAndServe(fmt.Sprintf(":%s", gae.PORT), indexHandler))
}
