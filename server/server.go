package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/kataras/iris"

	"github.com/YuShuanHsieh/trello-transform/models"
	"github.com/YuShuanHsieh/trello-transform/transform"
)

type Server struct {
	web           *iris.Application
	configuration *ServerConfiguration
}

func Default() *Server {
	config := defaultConfiguration()
	return New(config)
}

func New(config *ServerConfiguration) *Server {
	var s Server
	s.web = iris.Default()
	s.configuration = config
	return &s
}

func (s *Server) Run() {
	s.transformData()
	s.web.Run(iris.Addr(fmt.Sprintf(":%d", s.configuration.Port)))
}

// Example for transform data
func (s *Server) transformData() {
	s.web.Post("/transform", func(ctx iris.Context) {
		file, header, err := ctx.FormFile("file")
		if err != nil {
			dispatchError(fmt.Errorf("Get file [%s] error: [%s]", header.Filename, err.Error()), ctx)
			return
		}
		content, err := ioutil.ReadAll(file)
		if err != nil {
			dispatchError(fmt.Errorf("Read file [%s] error: [%s]", header.Filename, err.Error()), ctx)
			return
		}

		tr := transform.New(content)
		tr.Selector(tr.SelectByList(&models.List{Name: "2019/01"}))
		tr.Use("list", transform.CardBriefFunc)
		tr.Use("reference", transform.ExtractReferenceFunc)
		tr.Use("label", transform.CountLabelsFunc)
		tr.Exec()

		json, _ := json.Marshal(tr.GetAllResult())
		ctx.StatusCode(iris.StatusOK)
		ctx.WriteString(string(json))
	})
}

func dispatchError(err error, ctx iris.Context) {
	res := make(map[string]string)
	res["message"] = fmt.Sprintf("%s", err.Error())
	ctx.StatusCode(iris.StatusInternalServerError)

	r, _ := json.Marshal(res)
	ctx.WriteString(string(r))
	ctx.StopExecution()
}
