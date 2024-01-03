package hello

import (
	"context"
	"fmt"
	"gfdemo/api/hello/v1"
	"gfdemo/internal/model"
	"gfdemo/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"
	"plugin"
)

// 目的和需求：部分go的核心文件不开源，例如验证，主程序核心逻辑等等

func init() {

	//方法一：直接执行二进制静态文件（不满足需求）
	//cmd := exec.Command("./valid")
	//// 将标准输出连接到当前程序的标准输出
	//cmd.Stdout = os.Stdout
	//// 运行命令
	//err := cmd.Run()
	//if err != nil {
	//	fmt.Println("Error:", err.Error())
	//}

	//方法二：执行插件库，调用动态编译文件
	contents := gfile.GetContents("cert.key")

	plug, err := plugin.Open("core/core.so")
	if err != nil {
		fmt.Println("Error loading plugin:", err)
		return
	}

	validSymbol, err := plug.Lookup("Valid")
	if err != nil {
		fmt.Println("Error Lookup plugin:", err)
		return
	}

	validFunc, ok := validSymbol.(func(string) string)
	if !ok {
		fmt.Println("Error asserting function type")
		return
	}
	result := validFunc(contents)
	fmt.Println(result)
}

func (c *ControllerV1) School(ctx context.Context, req *v1.SchoolReq) (res *v1.SchoolRes, err error) {

	getRes, err := service.School().GetOne(ctx)
	g.Dump(gconv.Map(getRes)["schoolName"])
	service.View().Render(ctx, model.View{
		ContentType: req.Type,
		Data:        gconv.Map(getRes)["schoolName"],
		Title:       "首页",
	})

	return
}
