package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

const usage = `
Usage: serve [-addr HOST:PORT]
Serve content for www.princebot.com.
`

var addr = flag.String(
	"addr", "",
	"HTTP service address (defaults to ${SITE_ADDR} if set or ':8080')",
)

func init() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, usage)
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr)
	}
}

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	flag.Parse()
	if *addr == "" {
		*addr = defaultAddr()
	}
	log.Fatal(serve(*addr))
}

func defaultAddr() string {
	if s := os.Getenv("SITE_ADDR"); s != "" {
		return s
	}
	return ":8080"
}

func serve(addr string) error {
	app := gin.Default()
	app.StaticFile("/", "site/index.html")
	app.StaticFile("/favicon.ico", "site/img/favicon.ico")
	app.StaticFile("/resume", "site/pdf/resume.pdf")
	app.StaticFS("/img", gin.Dir("site/img", false))
	app.StaticFS("/css", gin.Dir("site/css", false))
	return app.Run(addr)
}
