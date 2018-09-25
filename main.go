package main

import (
  . "fmt"
  "github.com/gin-gonic/gin"
  "myteam/router"
  "myteam/model"
)

func main() {
  run(":8080")
}

func run(PORT string) {
  var server *gin.Engine = gin.Default()
  
  model.Init()
  
  router.Register(server)

  Printf("App running on %v\n", PORT)
  server.Run(PORT)
}
