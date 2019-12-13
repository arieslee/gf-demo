package action

import (
	"fmt"
	"gf-app/app/service"
	"github.com/gogf/gf/net/ghttp"
)

type CategoryAction struct {

}

func NewCategoryAction() *CategoryAction {
	return &CategoryAction{}
}

func (ca *CategoryAction) Update(r *ghttp.Request, id int) {
	cateService := service.NewCategoryService()
	_, err := cateService.Update(r,id)
	if err != nil{
		fmt.Println(err.Error())
	}
	//r.Response.RedirectTo("/admin/post/category")
}