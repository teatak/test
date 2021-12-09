package main

import (
	"os"

	"github.com/teatak/cart"
)

func main() {
	r := cart.Default()
	hostName, _ := os.Hostname()
	r.Use("/", func(c *cart.Context, n cart.Next) {
		c.JSON(200, cart.H{
			"code":   200,
			"path":   c.Request.RequestURI,
			"server": hostName,
			"host":   c.Request.URL,
		})
		n()
	})
	r.Run()
}
