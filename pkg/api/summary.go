// @BeeOverwrite YES
// @BeeGenerateTime 20201012_230803
package api

import (
	"github.com/spf13/cast"
	"shorturl-admin/pkg/model"
	"shorturl-admin/pkg/mus"
	"shorturl-admin/pkg/router/core"
	"shorturl-admin/pkg/trans"
)

func SummaryList(c *core.Context) {
	req := &trans.ReqPage{}
	if err := c.Bind(req); err != nil {
		c.JSONErrTips("参数错误", err)
		return
	}

	query := model.Conds{}

	if v := c.Query("id"); v != "" {
		query["id"] = v
	}

	if v := c.Query("date"); v != "" {
		query["date"] = v
	}

	if v := c.Query("author"); v != "" {
		query["author"] = v
	}

	total, list := model.SummaryListPage(query, req)
	c.JSONList(list, req.Current, req.PageSize, total)
}

func SummaryInfo(c *core.Context) {
	reqId := cast.ToInt(c.Query("id"))
	if reqId == 0 {
		c.JSONErrTips("request is error", nil)
		return
	}

	info, _ := model.SummaryInfo(mus.Db, reqId)

	c.JSONOK(info)
}

func SummaryCreate(c *core.Context) {
	req := &model.Summary{}
	if err := c.Bind(req); err != nil {
		c.JSONErrTips("参数错误", err)
		return
	}

	err := model.SummaryCreate(mus.Db, req)
	if err != nil {
		c.JSONErrTips("创建失败", err)
		return
	}
	c.JSONOK(req)
}

func SummaryDelete(c *core.Context) {
	reqJson := make(map[string]interface{}, 0)
	err := c.Bind(&reqJson)
	if err != nil {
		c.JSONErrTips("request is error: "+err.Error(), err)
		return
	}

	id := cast.ToInt(reqJson["id"])
	if id == 0 {
		c.JSONErrTips("id is error: ", nil)
		return
	}

	err = model.SummaryDelete(mus.Db, id)
	if err != nil {
		c.JSONErrTips("删除失败", err)
		return
	}
	c.JSONOK()
}

func SummaryUpdate(c *core.Context) {
	reqJson := make(map[string]interface{}, 0)
	err := c.Bind(&reqJson)
	if err != nil {
		c.JSONErrTips("request is error: "+err.Error(), err)
		return
	}

	id := cast.ToInt(reqJson["id"])
	if id == 0 {
		c.JSONErrTips("id is error: ", nil)
		return
	}

	err = model.SummaryUpdate(mus.Db, id, model.Ups{

		"id": reqJson["id"],

		"date": reqJson["date"],

		"author": reqJson["author"],
	})
	if err != nil {
		c.JSONErrTips("更新失败", err)
		return
	}
	c.JSONOK()
}
