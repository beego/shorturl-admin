// @BeeOverwrite YES
// @BeeGenerateTime 20201012_230803
package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RespList struct {
	Data     interface{} `json:"data"`
	Current  int         `json:"current"`
	PageSize int         `json:"pageSize"`
	Total    int         `json:"total"`
}

const (
	MsgOk  = 0
	MsgErr = 1
)

type HandlerFunc func(c *Context)

func Handle(h HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := &Context{
			c,
		}
		h(ctx)
	}
}

type Context struct {
	*gin.Context
}

// JSONResult json
type JSONResult struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func (c *Context) JSONOK(result ...interface{}) {
	j := new(JSONResult)
	j.Code = 0
	j.Message = "成功"
	if len(result) > 0 {
		j.Data = result[0]
	} else {
		j.Data = ""
	}
	c.JSON(http.StatusOK, j)
	return
}

func (c *Context) JSONErrTips(msg string, err error) {
	result := new(JSONResult)
	result.Code = MsgErr
	if err != nil {
		fmt.Println("info is", result.Message, "============== err is", err.Error())
	}
	result.Message = msg
	c.JSON(http.StatusOK, result)
	return
}

func (c *Context) JSONList(data interface{}, current, pageSize, total int) {
	j := new(RespList)
	j.Data = data
	j.Current = current
	j.PageSize = pageSize
	j.Total = total
	c.JSON(http.StatusOK, j)
	return
}
