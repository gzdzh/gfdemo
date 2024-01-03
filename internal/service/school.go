// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	ISchool interface {
		GetOne(ctx context.Context) (data interface{}, err error)
	}
)

var (
	localSchool ISchool
)

func School() ISchool {
	if localSchool == nil {
		panic("implement not found for interface ISchool, forgot register?")
	}
	return localSchool
}

func RegisterSchool(i ISchool) {
	localSchool = i
}
