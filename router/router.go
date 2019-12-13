package router

import (
    "gf-app/app/api/hello"
    "gf-app/app/module/admin/action"
    "github.com/gogf/gf/frame/g"
    "github.com/gogf/gf/net/ghttp"
)

// 统一路由注册.
func init() {
    g.Server().BindHandler("/", hello.Handler)
    s :=g.Server()
    s.Group("/admin", func(r *ghttp.RouterGroup) {
        r.POST("/post/category/edit/:id", func(r *ghttp.Request) {
            pc := action.NewCategoryAction()
            id := r.GetInt("id")
            pc.Update(r, id)
        })
    })
}
