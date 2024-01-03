package view

import (
	"context"
	"gfdemo/internal/model"
	"gfdemo/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gmode"
)

type sView struct{}

func init() {
	service.RegisterView(&sView{})
}

// RenderTpl 渲染指定模板页面
func (s *sView) RenderTpl(ctx context.Context, tpl string, data ...model.View) {
	var (
		viewObj  = model.View{}
		viewData = make(g.Map)
		request  = g.RequestFromCtx(ctx)
	)
	if len(data) > 0 {
		viewObj = data[0]
	}
	if viewObj.Title == "" {
		viewObj.Title = g.Cfg().MustGet(ctx, `setting.title`).String()
	} else {
		viewObj.Title = viewObj.Title + ` - ` + g.Cfg().MustGet(ctx, `setting.title`).String()
	}
	if viewObj.Keywords == "" {
		viewObj.Keywords = g.Cfg().MustGet(ctx, `setting.keywords`).String()
	}
	if viewObj.Description == "" {
		viewObj.Description = g.Cfg().MustGet(ctx, `setting.description`).String()
	}
	// 去掉空数据
	viewData = gconv.Map(viewObj)
	for k, v := range viewData {
		if g.IsEmpty(v) {
			delete(viewData, k)
		}
	}
	// 内置对象
	viewData["BuildIn"] = &viewBuildIn{httpRequest: request}
	// 内容模板
	if viewData["MainTpl"] == nil {
		viewData["MainTpl"] = s.getDefaultMainTpl(ctx)
	}
 
	// 渲染模板
	_ = request.Response.WriteTpl(tpl, viewData)
	// 开发模式下，在页面最下面打印所有的模板变量
	if gmode.IsDevelop() {
		err := request.Response.WriteTplContent(`{{dump .}}`, viewData)
		if err != nil {
			panic(err)
		}
	}
}

// Render 渲染默认模板页面
func (s *sView) Render(ctx context.Context, data ...model.View) {
	s.RenderTpl(ctx, g.Cfg().MustGet(ctx, "viewer.indexLayout").String(), data...)
}

// 获取视图存储目录
func (s *sView) getViewFolderName(ctx context.Context) string {
	return gstr.Split(g.Cfg().MustGet(ctx, "viewer.indexLayout").String(), "/")[0]
}

// 获取自动设置的MainTpl
func (s *sView) getDefaultMainTpl(ctx context.Context) string {
	var (
		viewFolderPrefix = s.getViewFolderName(ctx)
		urlPathArray     = gstr.SplitAndTrim(g.RequestFromCtx(ctx).URL.Path, "/")
		mainTpl          string
	)
	if len(urlPathArray) > 0 && urlPathArray[0] == viewFolderPrefix {
		urlPathArray = urlPathArray[1:]
	}
	switch {
	case len(urlPathArray) == 2:
		// 如果2级路由为数字，那么为模块的详情页面，那么路由固定为/xxx/detail。
		// 如果需要定制化内容模板，请在具体路由方法中设置MainTpl。
		if gstr.IsNumeric(urlPathArray[1]) {
			urlPathArray[1] = "detail"
		}
		mainTpl = viewFolderPrefix + "/" + gfile.Join(urlPathArray[0], urlPathArray[1]) + ".html"
	case len(urlPathArray) == 1:
		mainTpl = viewFolderPrefix + "/" + urlPathArray[0] + "/index.html"
	default:
		// 默认首页内容
		mainTpl = viewFolderPrefix + "/index/index.html"
	}
	return gstr.TrimLeft(mainTpl, "/")
}
