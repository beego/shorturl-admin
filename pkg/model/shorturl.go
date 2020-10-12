// @BeeOverwrite YES
// @BeeGenerateTime 20201012_230803
package model

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"shorturl-admin/pkg/mus"
	"shorturl-admin/pkg/trans"
)

type Shorturl struct {
	Id        int    `gorm:"auto"json:"id" form:"id"`            // id
	Title     string `gorm:"size(255)"json:"title" form:"title"` // 标题
	SummaryId int    `json:"summaryId" form:"summaryId"`         // 汇总id
	Url       string `gorm:"size(255)"json:"url" form:"url"`     // url

}

type Shorturls []*Shorturl

func (t *Shorturl) TableName() string {
	return "shorturl"
}

// AddShorturl insert a new Shorturl into database and returns
// last inserted Id on success.
func ShorturlCreate(db *gorm.DB, data *Shorturl) (err error) {

	if err = db.Create(data).Error; err != nil {
		mus.Logger.Error("create shorturl error", zap.Error(err))
		return
	}
	return
}

func ShorturlUpdate(db *gorm.DB, paramId int, ups Ups) (err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}

	if err = db.Table("shorturl").Where(sql, binds...).Updates(ups).Error; err != nil {
		mus.Logger.Error("shorturl update error", zap.Error(err))
		return
	}
	return
}

// UpdateX Update的扩展方法，根据Cond更新一条或多条记录
func ShorturlUpdateX(db *gorm.DB, conds Conds, ups Ups) (err error) {
	sql, binds := BuildQuery(conds)
	if err = db.Table("shorturl").Where(sql, binds...).Updates(ups).Error; err != nil {
		mus.Logger.Error("shorturl update error", zap.Error(err))
		return
	}
	return
}

// Delete 根据主键删除一条记录。如果有delete_time则软删除，否则硬删除。
func ShorturlDelete(db *gorm.DB, paramId int) (err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}

	if err = db.Table("shorturl").Where(sql, binds...).Delete(&Shorturl{}).Error; err != nil {
		mus.Logger.Error("shorturl delete error", zap.Error(err))
		return
	}

	return
}

// DeleteX Delete的扩展方法，根据Cond删除一条或多条记录。如果有delete_time则软删除，否则硬删除。
func ShorturlDeleteX(db *gorm.DB, conds Conds) (err error) {
	sql, binds := BuildQuery(conds)

	if err = db.Table("shorturl").Where(sql, binds...).Delete(&Shorturl{}).Error; err != nil {
		mus.Logger.Error("shorturl delete error", zap.Error(err))
		return
	}

	return
}

// Info 根据PRI查询单条记录
func ShorturlInfo(db *gorm.DB, paramId int) (resp Shorturl, err error) {

	var sql = "`id`= ?"

	var binds = []interface{}{paramId}

	if err = db.Table("shorturl").Where(sql, binds...).First(&resp).Error; err != nil {
		mus.Logger.Error("shorturl info error", zap.Error(err))
		return
	}
	return
}

// InfoX Info的扩展方法，根据Cond查询单条记录
func ShorturlInfoX(db *gorm.DB, conds Conds) (resp Shorturl, err error) {

	sql, binds := BuildQuery(conds)

	if err = db.Table("shorturl").Where(sql, binds...).First(&resp).Error; err != nil {
		mus.Logger.Error("shorturl info error", zap.Error(err))
		return
	}
	return
}

// List 查询list，extra[0]为sorts
func ShorturlList(conds Conds, extra ...string) (resp []*Shorturl, err error) {

	sql, binds := BuildQuery(conds)

	sorts := ""
	if len(extra) >= 1 {
		sorts = extra[0]
	}
	if err = mus.Db.Table("shorturl").Where(sql, binds...).Order(sorts).Find(&resp).Error; err != nil {
		mus.Logger.Error("shorturl info error", zap.Error(err))
		return
	}
	return
}

// ListMap 查询map，map遍历的时候是无序的，所以指定sorts参数没有意义
func ShorturlListMap(conds Conds) (resp map[int]*Shorturl, err error) {

	sql, binds := BuildQuery(conds)

	mysqlSlice := make([]*Shorturl, 0)
	resp = make(map[int]*Shorturl, 0)
	if err = mus.Db.Table("shorturl").Where(sql, binds...).Find(&mysqlSlice).Error; err != nil {
		mus.Logger.Error("shorturl info error", zap.Error(err))
		return
	}
	for _, value := range mysqlSlice {
		resp[value.Id] = value
	}
	return
}

// ListPage 根据分页条件查询list
func ShorturlListPage(conds Conds, reqList *trans.ReqPage) (total int, respList Shorturls) {
	respList = make(Shorturls, 0)

	if reqList.PageSize == 0 {
		reqList.PageSize = 10
	}
	if reqList.Current == 0 {
		reqList.Current = 1
	}
	sql, binds := BuildQuery(conds)

	db := mus.Db.Table("shorturl").Where(sql, binds...)
	respList = make([]*Shorturl, 0)
	db.Count(&total)
	db.Order(reqList.Sort).Offset((reqList.Current - 1) * reqList.PageSize).Limit(reqList.PageSize).Find(&respList)
	return
}
