// This is auto-generated by gf cli tool. You may not really want to edit it.

package model

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// BgCategory is the golang structure for table bg_category.
type BgCategory struct {
    Id        int    `orm:"id,primary"       json:"id"`          
    CateName  string `orm:"cate_name,unique" json:"cate_name"`   
    Slug      string `orm:"slug,unique"      json:"slug"`        
    Counts    int    `orm:"counts"           json:"counts"`      
    ParentId  int    `orm:"parent_id"        json:"parent_id"`   
    Intro     string `orm:"intro"            json:"intro"`       
    ListOrder int    `orm:"list_order"       json:"list_order"`  
    CreatedAt int64    `orm:"created_at"       json:"created_at"`
    UpdatedAt int64    `orm:"updated_at"       json:"updated_at"`
    Cover     string `orm:"cover"            json:"cover"`       
    Template  string `orm:"template"         json:"template"`    
    Status    int    `orm:"status"           json:"status"`      
}

var (
	// TableBgCategory is the table name of bg_category.
	TableBgCategory = "bg_category"
	// ModelBgCategory is the model object of bg_category.
	ModelBgCategory = g.DB("default").Table(TableBgCategory).Safe()
)

// Inserts does "INSERT...INTO..." statement for inserting current object into table.
func (r *BgCategory) Insert() (result sql.Result, err error) {
	return ModelBgCategory.Data(r).Insert()
}

// Replace does "REPLACE...INTO..." statement for inserting current object into table.
// If there's already another same record in the table (it checks using primary key or unique index),
// it deletes it and insert this one.
func (r *BgCategory) Replace() (result sql.Result, err error) {
	return ModelBgCategory.Data(r).Replace()
}

// Save does "INSERT...INTO..." statement for inserting/updating current object into table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *BgCategory) Save() (result sql.Result, err error) {
	return ModelBgCategory.Data(r).Save()
}

// Update does "UPDATE...WHERE..." statement for updating current object from table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *BgCategory) Update() (result sql.Result, err error) {
	return ModelBgCategory.Data(r).Where(gdb.GetWhereConditionOfStruct(r)).Update()
}

// Delete does "DELETE FROM...WHERE..." statement for deleting current object from table.
func (r *BgCategory) Delete() (result sql.Result, err error) {
	return ModelBgCategory.Where(gdb.GetWhereConditionOfStruct(r)).Delete()
}