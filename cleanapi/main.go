package cleanapi

import (
  "github.com/gin-gonic/gin"
)


var ( portalController controllers.Portalcontroller = controllers.PortalControllerImpl())


func main(){
  r := gin.Default()
  portalRoutes := r.Group("user/")
   {
    portalRoutes.POST("/login", portalController.Login)
    portalRoutes.POST("/register", portalController.Register)
    portalRoutes.POST("/addUser", portalController.AddUser)
    portalRoutes.POST("/addubject", portalController.AddSubject)
  }

  r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
