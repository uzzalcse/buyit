package controllers

import (
	"buyit/dao"
	"encoding/json"
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
    // query := map[string]interface{}{
    //     "query": map[string]interface{}{
    //         "nested": map[string]interface{}{
    //             "path": "products",
    //             "query": map[string]interface{}{
    //                 "match_prefix": map[string]interface{}{
    //                     "products.product_name": queryParams,
    //                 },
    //             },
    //         },
    //     },
    // }

    query2 := map[string]interface{}{
        "query": map[string]interface{}{
            "match_phrase_prefix": map[string]interface{}{
                "products.product_name": map[string]interface{}{
                    "query": queryParams,
                    "max_expansions": 10,
                },
            },
        },
    }

	q, _ := json.Marshal(query2)
	fmt.Println(string(q))

	// Execute search using ESClient
	res, err := f.esClient.ExecuteSearch(query2)
	if err != nil {
		f.Data["json"] = map[string]string{"error": fmt.Sprintf("Failed to fetch flight details: %v", err)}
		f.ServeJSON()
		return
	}

	// Send response
	f.Data["json"] = res
	f.ServeJSON()
}