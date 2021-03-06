// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"g/app/dao/internal"
)

// tUserGameAuthDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type tUserGameAuthDao struct {
	*internal.TUserGameAuthDao
}

var (
	// TUserGameAuth is globally public accessible object for table t_user_game_auth operations.
	TUserGameAuth = tUserGameAuthDao{
		internal.NewTUserGameAuthDao(),
	}
)

// Fill with you ideas below.
