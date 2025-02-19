package controllers

import (
	"buyit/dao"
	"fmt"
	"strings"

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

	query := map[string]interface{}{
		"size": 20,  // Number of results to return
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"should": []map[string]interface{}{
					{
						"match_phrase_prefix": map[string]interface{}{
							"products.product_name": map[string]interface{}{
								"query": queryParams,
								"max_expansions": 50,
								"boost": 4,
							},
						},
					},
					{
						"match": map[string]interface{}{
							"products.product_name": map[string]interface{}{
								"query": queryParams,
								"operator": "or",
								"fuzziness": "AUTO",
								"prefix_length": 1,
								"boost": 2,
							},
						},
					},
					{
						"wildcard": map[string]interface{}{
							"products.product_name": map[string]interface{}{
								"value": fmt.Sprintf("*%s*", strings.ToLower(queryParams)),
								"boost": 1,
							},
						},
					},
				},
				"minimum_should_match": 1,
			},
		},
		"highlight": map[string]interface{}{
			"fields": map[string]interface{}{
				"products.product_name": map[string]interface{}{},
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
