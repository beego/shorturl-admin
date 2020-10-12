// @BeeOverwrite YES
// @BeeGenerateTime 20201012_230803
package api

type ResponseData struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
