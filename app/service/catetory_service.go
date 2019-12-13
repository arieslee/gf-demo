package service

import (
	"database/sql"
	"errors"
	"gf-app/app/model"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
)
const (
	CacheKey string = "categoryCache"
)
type CategoryService struct {
	Base *BaseService
}

func NewCategoryService() *CategoryService {
	return &CategoryService{
		Base:new(BaseService),
	}
}

func (cs *CategoryService) CountBy(where string, params []interface{}) (int, error) {
	count, err := model.ModelBgCategory.Where(where, params).Count()
	if err != nil{
		return 0,errors.New(err.Error())
	}
	return count,nil
}
func (cs *CategoryService) ListBy(r *ghttp.Request,where string, params []interface{}) (map[string]interface{}, int, error) {
	input := map[string]interface{}{
		"where":where,
		"params":params,
		"page":1,
	}
	res, err := cs.Base.List(r, model.TableBgCategory, input)
	count := gconv.Int(res["totalCount"])
	if count > 0{
		 /*cateList := make(map[int]interface{})
		 for _,value := range gconv.Map(res["list"]){
		 	cate := gconv.Map(value)
		 	cateId := gconv.Int(cate["id"])
		 	cateList[cateId] = value
		 }
		 res["list"] = cateList*/
	}
	return res,count,err
}
func (cs *CategoryService) FindBy(where string, params []interface{}) (*model.BgCategory, error){
	var (
		row *model.BgCategory
	)
	/*db := g.DB()
	err := db.Table(model.TableBgCategory).Where(where, params).Limit(1).Struct(&row)
	if err == sql.ErrNoRows{
		return row, errors.New("暂无内容")
	}*/
	input := map[string]interface{}{
		"where":where,
		"params":params,
	}
	ar := cs.Base.FindBy(model.TableBgCategory, input)
	err := ar.Struct(&row)
	return row, err
}
func (cs *CategoryService) FindById(id int) (*model.BgCategory, error){
	var (
		where string
		params []interface{}
	)
	where = "id=?"
	params = append(params,id)
	row,err := cs.FindBy(where, params)
	return row,err
}
func (cs *CategoryService) DeleteBy(where string, params []interface{}) (sql.Result,error) {
	db := g.DB()
	res,err := db.Table(model.TableBgCategory).Where(where, params).Delete()
	return res,err
}
func (cs *CategoryService) DeleteById(id int) (sql.Result,error){
	var (
		where string
		params []interface{}
	)
	where = "id=?"
	params = append(params,id)
	return cs.DeleteBy(where, params)
}
func (cs *CategoryService) Create(r *ghttp.Request) (*model.BgCategory, error) {
	request := r.GetPostMap()
	rules := map[string]string {
		"cateName"  : "required|length:1,200",
	}
	msgs  := map[string]interface{} {
		"cateName" : "标题不能为空|标题的长度应当在:1到:128之间",
	}
	post := &model.BgCategory{}
	err := r.GetPostToStruct(post)
	if err != nil{
		return post, errors.New(err.Error())
	}
	if e := gvalid.CheckMap(request, rules, msgs); e != nil {
		return post, errors.New(e.FirstString())
	}
	nowTime := gtime.Now().Unix()
	post.CreatedAt = int64(nowTime)
	_,insErr := post.Insert()
	if insErr != nil{
		return post, errors.New(insErr.Error())
	}
	return post, nil
}
func (cs *CategoryService) Update(r *ghttp.Request, id int)(*model.BgCategory, error) {
	request := r.GetPostMap()
	rules := map[string]string {
		"cateName"  : "required|length:1,200",
	}
	msgs  := map[string]interface{} {
		"cateName" : "分类名称不能为空|分类名称的长度应当在:1到:128之间",
	}
	post := &model.BgCategory{}
	if e:=gvalid.CheckMap(request, rules, msgs); e!= nil{
		return post, errors.New(e.FirstString())
	}
	_ = r.GetPostToStruct(post)
	nowTime := gtime.Now().Unix()
	post.Id = id
	post.UpdatedAt = int64(nowTime)
	db := g.DB()
	db.SetDebug(true)
	_,updateErr := post.Update()
	//_,updateErr := db.Table(model.TableBgCategory).Data(post).Where("id=?", []interface{}{id}).Update()
	if updateErr != nil{
		return post, errors.New(updateErr.Error())
	}
	var (
		where string
		params []interface{}
	)
	cateList, _, _ := cs.ListBy(r,where, params)
	cacheList := make(map[int]interface{})
	for _,value := range gconv.SliceAny(cateList["list"]){
		row := gconv.Map(value)
		id := gconv.Int(row["id"])
		cacheList[id] = row
	}
	c := gcache.New()
	// 设置缓存，不过期
	c.Set(CacheKey, cacheList, 0)
	// 关闭缓存对象，让GC回收资源
	c.Close()
	return post, nil
}
func (cs *CategoryService) GetCacheData() interface{} {
	c := gcache.New()
	data := c.Get(CacheKey)
	if data == nil{
		where := "id>?"
		params := []interface{}{
			0,
		}
		var cateList []*model.BgCategory
		db := g.DB()
		err := db.Table(model.TableBgCategory).Where(where, params).OrderBy("created_at DESC").Structs(&cateList)
		if err != nil{
			return nil
		}
		cacheList := make(map[int]interface{})
		for _,value := range cateList{
			row := gconv.Map(value)
			id := gconv.Int(row["id"])
			cacheList[id] = row
		}
		c.Set(CacheKey, cacheList, 0)
		data = cacheList
	}
	// 关闭缓存对象，让GC回收资源
	c.Close()
	return data
}