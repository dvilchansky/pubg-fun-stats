package controllers

import (
	"github.com/kataras/iris"
	"pubg-fun-stats/web/services"
	"sort"
)

type MatchController struct {
	Service services.MatchService
}

func (c *MatchController) Get(ctx iris.Context) {
	userName := ctx.Params().Get("name")
	if userName != "" {
		matches, err := c.Service.RequestPlayerMatches(userName, 10)
		if err != nil {
			_, _ = ctx.JSON(iris.Map{
				"data": matches,
			})
		}
		sort.Slice(matches, func(i, j int) bool {
			return matches[i].CreatedAt.After(matches[j].CreatedAt)
		})
		if err != nil {
			panic(err.Error())
		}
		ctx.JSON(iris.Map{
			"data": matches,
		})
	}
}
