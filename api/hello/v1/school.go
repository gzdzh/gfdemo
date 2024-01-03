package v1

import "github.com/gogf/gf/v2/frame/g"

type SchoolReq struct {
	g.Meta `path:"/school" tags:"School" method:"get" summary:"You first school api"`
	Type   string `json:"type"   in:"query" dc:"内容模型"`
}
type SchoolRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
