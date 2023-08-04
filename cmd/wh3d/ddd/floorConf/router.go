package floorConf

import (
	"github.com/kataras/iris/v12"
)

func Router(p iris.Party) {
	p.Get("/", apiList)

	p.Get("/pc", floorConfPageListApi)
	p.Put("/", floorConfEdit)
}
