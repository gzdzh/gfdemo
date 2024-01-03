package view

import (
	"fmt"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gmode"

	"gfdemo/internal/consts"
)

// 视图自定义方法管理对象
type viewBuildIn struct {
	httpRequest *ghttp.Request
}

// IsNew 最新文章，12小时内为最新文章
func (s *viewBuildIn) IsNew(gt *gtime.Time) bool {
	if gt == nil {
		return false
	}
	n := gtime.Now().Timestamp()
	t := gt.Timestamp()

	var hs int64 = 3600

	// 最新文章
	if n-t < hs*12 {
		return true
	}

	return false
}

// GenderFont 根据性别字段内容返回性别的font。
func (s *viewBuildIn) GenderFont(gender int) string {
	switch gender {
	case consts.UserGenderMale:
		return "&#xe651;"
	case consts.UserGenderFemale:
		return "&#xe636;"
	default:
		return "&#xead2;"
	}
}

// Gender 根据性别字段内容返回性别。
func (s *viewBuildIn) Gender(gender int) string {
	switch gender {
	case consts.UserGenderMale:
		return "男"
	case consts.UserGenderFemale:
		return "女"
	default:
		return "未知"
	}
}

// ContentTypeName 根据性别字段内容返回性别。
func (s *viewBuildIn) ContentTypeName(contentType string) string {
	switch contentType {
	case consts.ContentTypeArticle:
		return "文章"
	case consts.ContentTypeAsk:
		return "问答"
	case consts.ContentTypeTopic:
		return "主题"
	default:
		return "未知"
	}
}

// Page 创建分页HTML内容
func (s *viewBuildIn) Page(total, size int) string {
	page := s.httpRequest.GetPage(total, size)
	page.LinkStyle = "page-link"
	page.SpanStyle = "page-link"
	page.PrevPageTag = "«"
	page.NextPageTag = "»"
	content := page.PrevPage() + page.PageBar() + page.NextPage()
	content = gstr.ReplaceByMap(content, map[string]string{
		"<span":  "<li class=\"page-item disabled\"><span",
		"/span>": "/span></li>",
		"<a":     "<li class=\"page-item\"><a",
		"/a>":    "/a></li>",
	})
	return content
}

// UrlPath 获取当前页面的Url Path.
func (s *viewBuildIn) UrlPath() string {
	return s.httpRequest.URL.Path
}

// FormatTime 格式化时间
func (s *viewBuildIn) FormatTime(gt *gtime.Time) string {
	if gt == nil {
		return ""
	}
	n := gtime.Now().Timestamp()
	t := gt.Timestamp()

	var ys int64 = 31536000
	var ds int64 = 86400
	var hs int64 = 3600
	var ms int64 = 60
	var ss int64 = 1

	var rs string

	d := n - t
	switch {
	case d > ys:
		rs = fmt.Sprintf("%d年前", int(d/ys))
	case d > ds:
		rs = fmt.Sprintf("%d天前", int(d/ds))
	case d > hs:
		rs = fmt.Sprintf("%d小时前", int(d/hs))
	case d > ms:
		rs = fmt.Sprintf("%d分钟前", int(d/ms))
	case d > ss:
		rs = fmt.Sprintf("%d秒前", int(d/ss))
	default:
		rs = "刚刚"
	}

	return rs
}

// Version 随机数 开发环境时间戳，线上为前端版本号
func (s *viewBuildIn) Version() string {
	var rand string
	if gmode.IsDevelop() {
		rand = gconv.String(gtime.TimestampMilli())
	} else {
		rand = consts.Version
	}
	return rand
}
