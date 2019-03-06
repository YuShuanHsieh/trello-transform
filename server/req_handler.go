package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"syscall"

	"github.com/gin-gonic/gin"

	"github.com/YuShuanHsieh/trello-transform/errors"
	"github.com/YuShuanHsieh/trello-transform/helpers/validators"
	"github.com/YuShuanHsieh/trello-transform/transform"
	"github.com/YuShuanHsieh/trello-transform/transform/defaultHandler"
	"github.com/YuShuanHsieh/trello-transform/transform/selector"
	"github.com/YuShuanHsieh/trello-transform/trello"
)

func (s *Server) stopServerHandler(c *gin.Context) {
	if !validators.IsFromLocalHost(c.Request.RemoteAddr) {
		dispatchError(
			errors.NewFromStr("Invalid Operation"),
			http.StatusForbidden,
			c,
		)
	}
	s.stop <- syscall.SIGINT
	dispatchSuccess("success", c)
}

func (s *Server) transformHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		dispatchError(
			errors.NewFromFormat("Get file [%s] error: [%s]", file.Filename, err.Error()),
			http.StatusBadRequest,
			c,
		)
		return
	}
	f, _ := file.Open()
	content, err := ioutil.ReadAll(f)
	if err != nil {
		dispatchError(
			errors.NewFromFormat("Read file [%s] error: [%s]", file.Filename, err.Error()),
			http.StatusInternalServerError,
			c,
		)
		return
	}

	var list trello.List
	err = getListFromFormData(&list, c)
	if err != nil {
		dispatchError(
			errors.NewFromFormat("Read list error: [%s]", err.Error()),
			http.StatusInternalServerError,
			c,
		)
		return
	}

	tr := transform.New(content)
	tr.Select(selector.ByList(list))
	tr.Use("list", defaultHandler.CardBriefHandler)
	tr.Use("reference", defaultHandler.ExtractReferenceHandler)
	tr.Use("label", defaultHandler.CountLabelsHandler)
	tr.Exec()

	json, _ := json.Marshal(tr.GetAllResult())
	dispatchSuccess(string(json), c)
}

func dispatchSuccess(result string, c *gin.Context) {
	c.String(http.StatusOK, result)
}

func dispatchError(err error, statusCode int, c *gin.Context) {
	res := make(map[string]string)
	res["message"] = err.Error()

	c.AbortWithStatusJSON(statusCode, res)
}

func getListFromFormData(list *trello.List, c *gin.Context) error {
	raw := c.PostForm("list")
	err := json.Unmarshal([]byte(raw), list)
	if err != nil {
		return err
	}
	return nil
}
