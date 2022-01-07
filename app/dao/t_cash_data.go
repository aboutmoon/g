// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/aboutmoon/g/app/dao/internal"
)

// tCashDataDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type tCashDataDao struct {
	*internal.TCashDataDao
}

var (
	// TCashData is globally public accessible object for table t_cash_data operations.
	TCashData = tCashDataDao{
		internal.NewTCashDataDao(),
	}
)

// Fill with you ideas below.
