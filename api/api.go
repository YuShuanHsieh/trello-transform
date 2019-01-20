package api

import (
	"github.com/kataras/iris"

	"github.com/YuShuanHsieh/trello-transform/config"
)

type API struct {
	app *iris.Application
}

func New() *API {
	var a API
	a.app = iris.Default()

	return &a
}

func (a *API) transformData() {
	a.app.Get("/transform", func(ctx iris.Context) {
		ctx.UploadFormFiles(config.FilePath)
	})
}
