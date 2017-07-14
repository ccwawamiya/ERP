package models

import "time"
import "github.com/astaxie/beego/orm"

// ModuleTable 模块(表)名称
type ModuleTable struct {
	ID           int64           `orm:"column(id);pk;auto" json:"id" form:"recordID"` //主键
	CreateUserID int64           `orm:"column(create_user_id);null" json:"-"`         //创建者
	UpdateUserID int64           `orm:"column(update_user_id);null" json:"-"`         //最后更新者
	CreateDate   time.Time       `orm:"auto_now_add;type(datetime)" json:"-"`         //创建时间
	UpdateDate   time.Time       `orm:"auto_now;type(datetime)" json:"-"`             //最后更新时间
	Name         string          `orm:"unique;size(50)" json:"name" form:"Name"`      //表名称
	Category     *ModuleCategory `orm:"rel(fk);null"`                                 //模块分类
}

func init() {
	orm.RegisterModel(new(ModuleTable))
}

// AddModuleTable insert a new ModuleTable into database and returns
// last inserted Id on success.
func AddModuleTable(m *ModuleTable, ormObj orm.Ormer) (id int64, err error) {
	id, err = ormObj.Insert(m)
	return
}