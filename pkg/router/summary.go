// @BeeOverwrite YES
// @BeeGenerateTime 20201012_230803
package router

import (
	"github.com/gin-gonic/gin"
	"shorturl-admin/pkg/api"
	"shorturl-admin/pkg/router/core"
)

func InitSummary(r gin.IRoutes) {
	core.RegisterUrl(r, "get", "/api/adminß/summary/info", api.SummaryInfo)
	core.RegisterUrl(r, "get", "/api/adminß/summary/list", api.SummaryList)
	core.RegisterUrl(r, "post", "/api/adminß/summary/create", api.SummaryCreate)
	core.RegisterUrl(r, "post", "/api/adminß/summary/update", api.SummaryUpdate)
	core.RegisterUrl(r, "post", "/api/adminß/summary/delete", api.SummaryDelete)
}
