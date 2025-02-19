package controllers

import (
	"buyit/dao"
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
	beecontext "github.com/beego/beego/v2/server/web/context"
)

type MainController struct {
	beego.Controller
	esClient *dao.ESClient
}

func (c *MainController) Get() {
	c.Data["Website"] = ""
	c.Data["Email"] = "uzzal.cse42@gmail.com"
	c.TplName = "index.html"
}

func (f *MainController) Init(ctx *beecontext.Context, controllerName, actionName string, app interface{}) {
	f.Controller.Init(ctx, controllerName, actionName, app)
	f.esClient = dao.Client
}

func (f *MainController) GetProducts() {
	queryParams := f.GetString("q")
	fmt.Println(queryParams)

    query := map[string]interface{}{
        "size": 20,
        "query": map[string]interface{}{
            "match": map[string]interface{}{
                "products.product_name": map[string]interface{}{
                    "query": queryParams,
                    "fuzziness": "AUTO",
                },
            },
        },
    }

	// Execute search using ESClient
	res, err := f.esClient.ExecuteSearch(query)
	if err != nil {
		f.Data["json"] = map[string]string{"error": fmt.Sprintf("Failed to fetch flight details: %v", err)}
		f.ServeJSON()
		return
	}

	// Send response
	f.Data["json"] = res
	f.ServeJSON()
}

func (f *MainController) GetProductDetailsByID() {
    id := f.Ctx.Input.Param(":id")
    fmt.Println("Searching for ID:", id)

    // Fetch document by ID instead of performing a search
    res, err := f.esClient.GetDocument("kibana_sample_data_ecommerce", id)
    if err != nil {
        f.Data["json"] = map[string]string{"error": fmt.Sprintf("Failed to fetch product details: %v", err)}
        f.ServeJSON()
        return
    }

    // Send response
    f.Data["json"] = res
    f.ServeJSON()
}
