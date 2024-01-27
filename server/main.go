package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hiagomf/gin-example/config"
	"github.com/hiagomf/gin-example/interface/http/user"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatal(err)
	}

	engine := gin.Default()

	// http://localhost:8080 + /usuarios

	user.Router(engine.Group("users"))

	engine.Run(fmt.Sprintf(":%d", config.Get().HTTPPort)) // listen and serve on 0.0.0.0:8080
}
