// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// ApDyqqUserDao is the manager for logic model data accessing and custom defined data operations functions management.
type ApDyqqUserDao struct {
	Table   string            // Table is the underlying table name of the DAO.
	Group   string            // Group is the database configuration group name of current DAO.
	Columns ApDyqqUserColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// ApDyqqUserColumns defines and stores column names for table ap_dyqq_user.
type ApDyqqUserColumns struct {
	Id                  string //
	Unionid             string // 微信唯一ID
	Uuid                string // 关联归因标识
	WxCode              string // 微信标识, 登录标识
	ZsOpenid            string // 游客登录,登录标识
	AppCode             string // 用户标识
	AppVersion          string // app版本
	WxOpenid            string // 微信openid
	ClientCode          string // 客户端标识码
	Vtype               string // 游戏版本
	Ip                  string //
	Mac                 string // mac地址
	AndroidId           string // 安卓设备唯一ID
	Imei                string // 国际移动设备识别码
	Oaid                string // 匿名设备标识符
	Channel             string // 包名
	Nickname            string // 用户昵称
	Sex                 string // 性别, [1:男,2:女]
	Avatar              string // 微信头像
	LuckyDraw           string // 抽奖次数
	CashSum             string // 总提现金额
	CashNum             string // 提现次数
	MusicCashSum        string // 总提现金额
	MusicCashNum        string // 总提现次数
	RedbagcoinSum       string // 红包币总数
	RedbagTimes         string // 普通领取次数
	RedbagDoubleTimes   string // 双倍奖励次数
	VideoCallbackNum    string // 视频回调次数
	ReceiveGiftTimes    string // 领取礼包数
	DistributeGiftTimes string // 下发礼包数
	SuccessTimes        string // 答题正确次数
	FailureTimes        string // 答题失败次数
	RetryTimes          string // 重复答题次数
	EcpmSum             string // 总ecpm
	EcpmMax             string // 最高ecpm
	EcpmAvg             string // 平均ecpm
	EcpmMin             string // 最低ecpm
	RedbagSum           string // 剩余红包数
	RedbagTotal         string // 总领取红包数
	IsTest              string // 是否为测试用户, 1:测试用户
	HadDelete           string // 是否已经被删除, 1:已删除;
	LastloginAt         string // 最后登录时间
	CreateAt            string // 创建时间
	VideoNum            string // 视频次数
	Level               string // 关卡数
	Diamond             string // 钻石
	LastChannel         string // 最后一次登录渠道
	LastAppCode         string // 最后一次登录appCode
	LastClientCode      string // 最后一次登录clientCode
	IsHacker            string // 是否是刷子
	IsDeleted           string // 是否删除；0:未删除；1:删除
	CreateTime          string // 创建时间
	UpdateTime          string // 更新时间
	IsReview            string // 是否是审核用户
	Idfa                string // 广告标识符
	Ua                  string // 用户代理
	RewardClickSum      string // 点击奖励总和
	RewardClickDaily    string // 当天点击奖励
	MusicCoin           string // 音符币
	MusicCoinTotal      string // 总领取音符币
	DeviceTags          string // 设备风险标签
	DeviceToken         string // 阿里云设备token
	RewardAdTimes       string // 点击奖励总和
	RewardGiftTimes     string // 点击奖励总和
	RedbagCashNum       string // 红包币提现总次数
	RedbagCashSum       string // 红包币提现总和
	Saving              string // 红包币存钱罐
	SavingMusicCoin     string // 音符币存钱罐
	Task                string // 任务
	Bonus               string // 新人福利
	Package             string // 包名
	Attribute           string // 是否归因
	Test                string // 字节测试
	IsSync              string // 是否同步
	Device              string // 设备号
	CardName            string // 身份证姓名
	CardNo              string // 身份证号
	Os                  string // 是否同步
	IgOpenid            string // 游客登录,登录标识
	IgUnionid           string // 游客登录,登录标识
	HackerTags          string // 拉黑标签
	UserTags            string // 用户标签
}

//  apDyqqUserColumns holds the columns for table ap_dyqq_user.
var apDyqqUserColumns = ApDyqqUserColumns{
	Id:                  "id",
	Unionid:             "unionid",
	Uuid:                "uuid",
	WxCode:              "wx_code",
	ZsOpenid:            "zs_openid",
	AppCode:             "app_code",
	AppVersion:          "app_version",
	WxOpenid:            "wx_openid",
	ClientCode:          "client_code",
	Vtype:               "vtype",
	Ip:                  "ip",
	Mac:                 "mac",
	AndroidId:           "android_id",
	Imei:                "imei",
	Oaid:                "oaid",
	Channel:             "channel",
	Nickname:            "nickname",
	Sex:                 "sex",
	Avatar:              "avatar",
	LuckyDraw:           "lucky_draw",
	CashSum:             "cash_sum",
	CashNum:             "cash_num",
	MusicCashSum:        "music_cash_sum",
	MusicCashNum:        "music_cash_num",
	RedbagcoinSum:       "redbagcoin_sum",
	RedbagTimes:         "redbag_times",
	RedbagDoubleTimes:   "redbag_double_times",
	VideoCallbackNum:    "video_callback_num",
	ReceiveGiftTimes:    "receive_gift_times",
	DistributeGiftTimes: "distribute_gift_times",
	SuccessTimes:        "success_times",
	FailureTimes:        "failure_times",
	RetryTimes:          "retry_times",
	EcpmSum:             "ecpm_sum",
	EcpmMax:             "ecpm_max",
	EcpmAvg:             "ecpm_avg",
	EcpmMin:             "ecpm_min",
	RedbagSum:           "redbag_sum",
	RedbagTotal:         "redbag_total",
	IsTest:              "is_test",
	HadDelete:           "had_delete",
	LastloginAt:         "lastlogin_at",
	CreateAt:            "create_at",
	VideoNum:            "video_num",
	Level:               "level",
	Diamond:             "diamond",
	LastChannel:         "last_channel",
	LastAppCode:         "last_app_code",
	LastClientCode:      "last_client_code",
	IsHacker:            "is_hacker",
	IsDeleted:           "is_deleted",
	CreateTime:          "create_time",
	UpdateTime:          "update_time",
	IsReview:            "is_review",
	Idfa:                "idfa",
	Ua:                  "ua",
	RewardClickSum:      "reward_click_sum",
	RewardClickDaily:    "reward_click_daily",
	MusicCoin:           "music_coin",
	MusicCoinTotal:      "music_coin_total",
	DeviceTags:          "device_tags",
	DeviceToken:         "device_token",
	RewardAdTimes:       "reward_ad_times",
	RewardGiftTimes:     "reward_gift_times",
	RedbagCashNum:       "redbag_cash_num",
	RedbagCashSum:       "redbag_cash_sum",
	Saving:              "saving",
	SavingMusicCoin:     "saving_music_coin",
	Task:                "task",
	Bonus:               "bonus",
	Package:             "package",
	Attribute:           "attribute",
	Test:                "test",
	IsSync:              "is_sync",
	Device:              "device",
	CardName:            "card_name",
	CardNo:              "card_no",
	Os:                  "os",
	IgOpenid:            "ig_openid",
	IgUnionid:           "ig_unionid",
	HackerTags:          "hacker_tags",
	UserTags:            "user_tags",
}

// NewApDyqqUserDao creates and returns a new DAO object for table data access.
func NewApDyqqUserDao() *ApDyqqUserDao {
	return &ApDyqqUserDao{
		Group:   "default",
		Table:   "ap_dyqq_user",
		Columns: apDyqqUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ApDyqqUserDao) DB() gdb.DB {
	return g.DB(dao.Group)
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ApDyqqUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ApDyqqUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
