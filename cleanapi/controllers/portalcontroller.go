package controllers

import(
"net/http"
"github.com/gin-gonic/gin"

)

type Portalcontroller  interface {
  Login(ctx *gin.Context)
  Register(ctx *gin.Context)
  AddUser(ctx *gin.Context)
  AddSubject(ctx *gin.Context)
}

type portalController struct{}

func PortalControllerImpl() Portalcontroller {

    return &portalController{}
}

func (c *portalController) Login(ctx *gin.Context){

  ctx.JSON(http.StatusOK, gin.H{
    "message": "You have succesfully logged in",
  })
}

func (c *portalController) Register(ctx *gin.Context){

  ctx.JSON(http.StatusOK, gin.H{
    "message": "You have succesfully Register",
  })
}

func (c *portalController) AddUser(ctx *gin.Context){

  ctx.JSON(http.StatusOK, gin.H{
    "message": "You have succesfully added a user",
  })
}

func (c *portalController) AddSubject(ctx *gin.Context){

  ctx.JSON(http.StatusOK, gin.H{
    "message": "You have succesfully added a subject",
  })
}
