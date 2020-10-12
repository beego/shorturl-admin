// @BeeOverwrite YES
// @BeeGenerateTime 20201012_230803
package router

import (
	"github.com/gin-gonic/gin"
	"shorturl-admin/pkg/api"
	"shorturl-admin/pkg/router/core"
)

func InitShorturl(r gin.IRoutes) {
	core.RegisterUrl(r, "get", "/api/adminß/shorturl/info", api.ShorturlInfo)
	core.RegisterUrl(r, "get", "/api/adminß/shorturl/list", api.ShorturlList)
	core.RegisterUrl(r, "post", "/api/adminß/shorturl/create", api.ShorturlCreate)
	core.RegisterUrl(r, "post", "/api/adminß/shorturl/update", api.ShorturlUpdate)
	core.RegisterUrl(r, "post", "/api/adminß/shorturl/delete", api.ShorturlDelete)
}
