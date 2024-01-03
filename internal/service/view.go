// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gfdemo/internal/model"
)

type (
	IView interface {
		// RenderTpl 渲染指定模板页面
		RenderTpl(ctx context.Context, tpl string, data ...model.View)
		// Render 渲染默认模板页面
		Render(ctx context.Context, data ...model.View)
	}
)

var (
	localView IView
)

func View() IView {
	if localView == nil {
		panic("implement not found for interface IView, forgot register?")
	}
	return localView
}

func RegisterView(i IView) {
	localView = i
}
