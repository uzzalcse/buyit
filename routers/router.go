package routers

import (
	"buyit/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/api/products/search", &controllers.MainController{}, "get:GetProducts")
	beego.Router("/api/products/:id", &controllers.MainController{}, "get:GetProductDetailsByID")
}
