package main

import (
	"flag"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	var port = flag.Uint("p", 8080, "server port")
	flag.Parse()

	r := gin.Default()
	r.Static("/", "./")
	r.Run("0.0.0.0:" + strconv.Itoa(int(*port))) // listen and serve on 0.0.0.0:8080
}
