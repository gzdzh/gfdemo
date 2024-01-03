// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// School is the golang structure for table school.
type School struct {
	Id         uint64      `json:"id"         ` //
	CreateTime *gtime.Time `json:"createTime" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" ` // 更新时间
	SchoolName string      `json:"schoolName" ` // 学校名称
}
