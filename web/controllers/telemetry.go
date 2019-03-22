package controllers

import (
	"github.com/kataras/iris"
	"pubg-fun-stats/web/services"
)

type TelemetryController struct {
	Service services.TelemetryService
}

func (c *TelemetryController) Get(ctx iris.Context) {
	urlParams := ctx.URLParams()
	endpointURL, ok := urlParams["endpointURL"]
	if ok {
		t, err := c.Service.RequestMatchTelemetry(endpointURL)
		if err != nil {
			panic(err.Error())
		}
		ctx.JSON(iris.Map{
			"data": t,
		})
	}
}
