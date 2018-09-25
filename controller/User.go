package controller

import (
  "github.com/gin-gonic/gin"
  "myteam/model"
  "myteam/types"
  . "fmt"
)

func UserIndex(c *gin.Context) {
  users, err := model.GetUsers()

  if err != nil {
    Println(err.Error())
  }

  c.JSON(200, users)
}

func UserShow(c *gin.Context) {
  user, err := model.GetUserByID(c.Param("id"))

  if err != nil {
    c.JSON(400, gin.H{
      "message": "Get user failed",
      "error": err.Error(),
    })
    return
  }

  c.JSON(200, gin.H{
    "message": "Get user success",
    "user": gin.H{
      "username": user.Username,
    },
  })
}

func UserStore(c *gin.Context) {
  var err error
  var user types.User

  c.BindJSON(&user)

  user, err = model.StoreUser(user)

  if err != nil {
    c.JSON(400, gin.H{
      "message": "Create user failed",
      "error": err.Error(),
    })
    return
  }

  c.JSON(200, gin.H{
    "message": "Create user success",
    "user": gin.H{
      "username": user.Username,
    },
  })
}

func UserUpdate(c *gin.Context) {
  var data types.User
  c.BindJSON(&data)

  user, err := model.UpdateUser(c.Param("id"), data)
  if err != nil {
    c.JSON(400, gin.H{
      "message": "Update user failed",
      "error": err.Error(),
    })
    return
  }

  c.JSON(200, gin.H{
    "message": "Update user success",
    "user": gin.H{
      "username": user.Username,
    },
  })
}

func UserDelete(c *gin.Context) {
  user, err := model.DeleteUser(c.Param("id"))

  if err != nil {
    c.JSON(400, gin.H{
      "message": "Delete user failed",
      "error": err.Error(),
    })
    return
  }

  c.JSON(200, gin.H{
    "message": "Delete user success",
    "user": gin.H{
      "username": user.Username,
    },
  })
}
