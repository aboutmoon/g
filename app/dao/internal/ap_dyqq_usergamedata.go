// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// ApDyqqUsergamedataDao is the manager for logic model data accessing and custom defined data operations functions management.
type ApDyqqUsergamedataDao struct {
	Table   string                    // Table is the underlying table name of the DAO.
	Group   string                    // Group is the database configuration group name of current DAO.
	Columns ApDyqqUsergamedataColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// ApDyqqUsergamedataColumns defines and stores column names for table ap_dyqq_usergamedata.
type ApDyqqUsergamedataColumns struct {
	Id            string // 主键
	Uid           string // 用户id
	ModelName     string // 模块名称
	ModelContents string // 模块内容
	UpdateTime    string // 更新时间
	CreateTime    string // 插入时间
	ClientTime    string // 客户端时间
	IsDeleted     string // 是否删除；0:未删除；1:删除
}

//  apDyqqUsergamedataColumns holds the columns for table ap_dyqq_usergamedata.
var apDyqqUsergamedataColumns = ApDyqqUsergamedataColumns{
	Id:            "id",
	Uid:           "uid",
	ModelName:     "model_name",
	ModelContents: "model_contents",
	UpdateTime:    "update_time",
	CreateTime:    "create_time",
	ClientTime:    "client_time",
	IsDeleted:     "is_deleted",
}

// NewApDyqqUsergamedataDao creates and returns a new DAO object for table data access.
func NewApDyqqUsergamedataDao() *ApDyqqUsergamedataDao {
	return &ApDyqqUsergamedataDao{
		Group:   "default",
		Table:   "ap_dyqq_usergamedata",
		Columns: apDyqqUsergamedataColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ApDyqqUsergamedataDao) DB() gdb.DB {
	return g.DB(dao.Group)
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ApDyqqUsergamedataDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ApDyqqUsergamedataDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}