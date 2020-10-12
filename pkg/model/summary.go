// @BeeOverwrite YES
// @BeeGenerateTime 20201012_230803
package model

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"shorturl-admin/pkg/mus"
	"shorturl-admin/pkg/trans"
)

type Summary struct {
	Id     int    `gorm:"auto"json:"id" form:"id"`              // id
	Date   string `gorm:"size(255)"json:"date" form:"date"`     // 标题
	Author string `gorm:"size(255)"json:"author" form:"author"` // 编辑

}

type Summarys []*Summary

func (t *Summary) TableName() string {
	return "summary"
}

// AddSummary insert a new Summary into database and returns
// last inserted Id on success.
func SummaryCreate(db *gorm.DB, data *Summary) (err error) {

	if err = db.Create(data).Error; err != nil {
		mus.Logger.Error("create summary error", zap.Error(err))
		return
	}
	return
}

func SummaryUpdate(db *gorm.DB, paramId int, ups Ups) (err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}

	if err = db.Table("summary").Where(sql, binds...).Updates(ups).Error; err != nil {
		mus.Logger.Error("summary update error", zap.Error(err))
		return
	}
	return
}

// UpdateX Update的扩展方法，根据Cond更新一条或多条记录
func SummaryUpdateX(db *gorm.DB, conds Conds, ups Ups) (err error) {
	sql, binds := BuildQuery(conds)
	if err = db.Table("summary").Where(sql, binds...).Updates(ups).Error; err != nil {
		mus.Logger.Error("summary update error", zap.Error(err))
		return
	}
	return
}

// Delete 根据主键删除一条记录。如果有delete_time则软删除，否则硬删除。
func SummaryDelete(db *gorm.DB, paramId int) (err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}

	if err = db.Table("summary").Where(sql, binds...).Delete(&Summary{}).Error; err != nil {
		mus.Logger.Error("summary delete error", zap.Error(err))
		return
	}

	return
}

// DeleteX Delete的扩展方法，根据Cond删除一条或多条记录。如果有delete_time则软删除，否则硬删除。
func SummaryDeleteX(db *gorm.DB, conds Conds) (err error) {
	sql, binds := BuildQuery(conds)

	if err = db.Table("summary").Where(sql, binds...).Delete(&Summary{}).Error; err != nil {
		mus.Logger.Error("summary delete error", zap.Error(err))
		return
	}

	return
}

// Info 根据PRI查询单条记录
func SummaryInfo(db *gorm.DB, paramId int) (resp Summary, err error) {

	var sql = "`id`= ?"

	var binds = []interface{}{paramId}

	if err = db.Table("summary").Where(sql, binds...).First(&resp).Error; err != nil {
		mus.Logger.Error("summary info error", zap.Error(err))
		return
	}
	return
}

// InfoX Info的扩展方法，根据Cond查询单条记录
func SummaryInfoX(db *gorm.DB, conds Conds) (resp Summary, err error) {

	sql, binds := BuildQuery(conds)

	if err = db.Table("summary").Where(sql, binds...).First(&resp).Error; err != nil {
		mus.Logger.Error("summary info error", zap.Error(err))
		return
	}
	return
}

// List 查询list，extra[0]为sorts
func SummaryList(conds Conds, extra ...string) (resp []*Summary, err error) {

	sql, binds := BuildQuery(conds)

	sorts := ""
	if len(extra) >= 1 {
		sorts = extra[0]
	}
	if err = mus.Db.Table("summary").Where(sql, binds...).Order(sorts).Find(&resp).Error; err != nil {
		mus.Logger.Error("summary info error", zap.Error(err))
		return
	}
	return
}

// ListMap 查询map，map遍历的时候是无序的，所以指定sorts参数没有意义
func SummaryListMap(conds Conds) (resp map[int]*Summary, err error) {

	sql, binds := BuildQuery(conds)

	mysqlSlice := make([]*Summary, 0)
	resp = make(map[int]*Summary, 0)
	if err = mus.Db.Table("summary").Where(sql, binds...).Find(&mysqlSlice).Error; err != nil {
		mus.Logger.Error("summary info error", zap.Error(err))
		return
	}
	for _, value := range mysqlSlice {
		resp[value.Id] = value
	}
	return
}

// ListPage 根据分页条件查询list
func SummaryListPage(conds Conds, reqList *trans.ReqPage) (total int, respList Summarys) {
	respList = make(Summarys, 0)

	if reqList.PageSize == 0 {
		reqList.PageSize = 10
	}
	if reqList.Current == 0 {
		reqList.Current = 1
	}
	sql, binds := BuildQuery(conds)

	db := mus.Db.Table("summary").Where(sql, binds...)
	respList = make([]*Summary, 0)
	db.Count(&total)
	db.Order(reqList.Sort).Offset((reqList.Current - 1) * reqList.PageSize).Limit(reqList.PageSize).Find(&respList)
	return
}
