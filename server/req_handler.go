package server

import (
	"encoding/json"
	"io/ioutil"
	"regexp"

	"github.com/kataras/iris"

	"github.com/YuShuanHsieh/trello-transform/errors"
	"github.com/YuShuanHsieh/trello-transform/models"
	"github.com/YuShuanHsieh/trello-transform/transform"
	"github.com/YuShuanHsieh/trello-transform/transform/selector"
)

func (s *Server) stopServerHandler(ctx iris.Context) {
	if !isFromLocalHost(ctx.Request().RemoteAddr) {
		dispatchError(errors.NewFromStr("Invalid Operation"), ctx)
	}
	// TODO should change to a specific function
	ctx.StatusCode(iris.StatusOK)
	ctx.WriteString("success")
	s.server.Shutdown(s.ctx)
}

func (s *Server) transformHandler(ctx iris.Context) {
	file, header, err := ctx.FormFile("file")
	if err != nil {
		dispatchError(errors.NewFromFormat("Get file [%s] error: [%s]", header.Filename, err.Error()), ctx)
		return
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		dispatchError(errors.NewFromFormat("Read file [%s] error: [%s]", header.Filename, err.Error()), ctx)
		return
	}

	tr := transform.New(content)
	tr.Select(selector.ByList(tr, models.List{Name: "2019/01"}))
	tr.Use("list", transform.CardBriefFunc)
	tr.Use("reference", transform.ExtractReferenceFunc)
	tr.Use("label", transform.CountLabelsFunc)
	tr.Exec()

	json, _ := json.Marshal(tr.GetAllResult())
	ctx.StatusCode(iris.StatusOK)
	ctx.WriteString(string(json))
}

func dispatchError(err error, ctx iris.Context) {
	res := make(map[string]string)
	res["message"] = err.Error()
	ctx.StatusCode(iris.StatusInternalServerError)

	r, _ := json.Marshal(res)
	ctx.WriteString(string(r))
	ctx.StopExecution()
}

func isFromLocalHost(remoteAddr string) bool {
	matched, _ := regexp.MatchString("^\\[::1\\]", remoteAddr)
	return matched
}
