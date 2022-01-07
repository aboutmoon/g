// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/aboutmoon/g/app/dao/internal"
)

// tGameJsonDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type tGameJsonDao struct {
	*internal.TGameJsonDao
}

var (
	// TGameJson is globally public accessible object for table t_game_json operations.
	TGameJson = tGameJsonDao{
		internal.NewTGameJsonDao(),
	}
)

// Fill with you ideas below.
