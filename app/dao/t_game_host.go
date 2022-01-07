// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"g/app/dao/internal"
)

// tGameHostDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type tGameHostDao struct {
	*internal.TGameHostDao
}

var (
	// TGameHost is globally public accessible object for table t_game_host operations.
	TGameHost = tGameHostDao{
		internal.NewTGameHostDao(),
	}
)

// Fill with you ideas below.
