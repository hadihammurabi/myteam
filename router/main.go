package router

import (
  "github.com/gin-gonic/gin"
  "myteam/controller"
  "myteam/middleware"
)

func Register(s *gin.Engine) {
  // Auth middleware
  s.Use(middleware.Auth())
  // Root Route
  s.GET("/", controller.Index)

  // Users Routes
  s.GET("/users", controller.UserIndex)
  s.GET("/users/:id", controller.UserShow)
  s.POST("/users", controller.UserStore)
  s.PUT("/users/:id", controller.UserUpdate)
  s.DELETE("/users/:id", controller.UserDelete)

  // Auth Route
  s.POST("/auth/signin", controller.AuthSignin)
}
