package controllers

import (
	"github.com/NyaaPantsu/manga/models"
	"github.com/NyaaPantsu/manga/utils/auth"
	"github.com/astaxie/beego"

	"encoding/json"
)

type AutoGenerated struct {
	LikeID  string `json:"like_id"`
	ComicID string `json:"comic_id"`
	Title   string `json:"title"`
	Image   string `json:"image"`
	URL     string `json:"url"`
}

// ImportController operations for Import
type ImportController struct {
	beego.Controller
}

// URLMapping ...
func (c *ImportController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetAll", c.GetAll)
}

// Post ...
// @Title Create
// @Description create Import
// @Param	body		body 	models.Import	true		"body for Import content"
// @Success 201 {object} models.Import
// @Failure 403 body is empty
// @router / [post]
func (c *ImportController) Post() {
	var v []AutoGenerated
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		for _, key := range v {
			name, _ := models.GetSeriesByName(key.Title)
			user, _ := models.GetUserByUsername(auth.GetUsername(c.Ctx))
			models.Follow(user.Id, name.Id)
		}
		c.Data["json"] = Response{
			Success: true,
		}

	} else {
		c.Data["json"] = Response{
			Success: false,
			Error:   err.Error(),
		}
	}
	c.ServeJSON()
}

// GetOne ...
// @Title GetOne
// @Description get Import by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Import
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ImportController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Import
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Import
// @Failure 403
// @router / [get]
func (c *ImportController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Import
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Import	true		"body for Import content"
// @Success 200 {object} models.Import
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ImportController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Import
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ImportController) Delete() {

}
