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

func ShorturlList(c *core.Context) {
	req := &trans.ReqPage{}
	if err := c.Bind(req); err != nil {
		c.JSONErrTips("参数错误", err)
		return
	}

	query := model.Conds{}

	if v := c.Query("id"); v != "" {
		query["id"] = v
	}

	if v := c.Query("title"); v != "" {
		query["title"] = v
	}

	if v := c.Query("summaryId"); v != "" {
		query["summaryId"] = v
	}

	if v := c.Query("url"); v != "" {
		query["url"] = v
	}

	total, list := model.ShorturlListPage(query, req)
	c.JSONList(list, req.Current, req.PageSize, total)
}

func ShorturlInfo(c *core.Context) {
	reqId := cast.ToInt(c.Query("id"))
	if reqId == 0 {
		c.JSONErrTips("request is error", nil)
		return
	}

	info, _ := model.ShorturlInfo(mus.Db, reqId)

	c.JSONOK(info)
}

func ShorturlCreate(c *core.Context) {
	req := &model.Shorturl{}
	if err := c.Bind(req); err != nil {
		c.JSONErrTips("参数错误", err)
		return
	}

	err := model.ShorturlCreate(mus.Db, req)
	if err != nil {
		c.JSONErrTips("创建失败", err)
		return
	}
	c.JSONOK(req)
}

func ShorturlDelete(c *core.Context) {
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

	err = model.ShorturlDelete(mus.Db, id)
	if err != nil {
		c.JSONErrTips("删除失败", err)
		return
	}
	c.JSONOK()
}

func ShorturlUpdate(c *core.Context) {
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

	err = model.ShorturlUpdate(mus.Db, id, model.Ups{

		"id": reqJson["id"],

		"title": reqJson["title"],

		"summary_id": reqJson["summaryId"],

		"url": reqJson["url"],
	})
	if err != nil {
		c.JSONErrTips("更新失败", err)
		return
	}
	c.JSONOK()
}
