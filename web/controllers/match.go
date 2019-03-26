package controllers

import (
	"github.com/kataras/iris"
	"github.com/labstack/gommon/log"
	"pubg-fun-stats/web/services"
	"sort"
)

type MatchController struct {
	Service services.MatchService
}

func (c *MatchController) Get(ctx iris.Context) {
	userName := ctx.Params().Get("name")
	if userName != "" {
		matches, err := c.Service.RequestPlayerMatches(userName, 1)
		if err != nil {
			log.Fatal(err.Error())
		}
		if len(matches) > 0 {
			sort.Slice(matches, func(i, j int) bool {
				return matches[i].CreatedAt.After(matches[j].CreatedAt)
			})
			_, err := ctx.JSON(iris.Map{
				"data": matches,
			})
			if err != nil {
				log.Fatal(err.Error())
			}
		} else {
			_, _ = ctx.JSON(iris.Map{
				"message": "No matches found for the given player name",
			})
		}
	}
}
