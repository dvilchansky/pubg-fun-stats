package controllers

import (
	"github.com/kataras/iris"
	"pubg-fun-stats/services"
)

type MatchController struct {
	Service services.MatchService
}

func (c *MatchController) Post(ctx iris.Context) {
	userName := ctx.Params().Get("name")
	if userName != "" {
		matches, err := c.Service.RequestPlayerMatches(userName)
		if err != nil {
			panic(err.Error())
		}
		ctx.JSON(iris.Map{
			"data": matches,
		})
	}
}
