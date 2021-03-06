// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"g/app/dao/internal"
)

// tAppsDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type tAppsDao struct {
	*internal.TAppsDao
}

var (
	// TApps is globally public accessible object for table t_apps operations.
	TApps = tAppsDao{
		internal.NewTAppsDao(),
	}
)

// Fill with you ideas below.
