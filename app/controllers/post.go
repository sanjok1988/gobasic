package controllers

import (
	"github.com/revel/revel"
	"webapp/app/models"
	"errors"
)

type Post struct {
	BaseController
}

func (c Post) Index() revel.Result {
	posts := []models.Post{}  // an array of Post model

	result := DB.Order("id desc").Find(&posts); //retrieve all the "posts" table records, and store them to the "posts" array
	err := result.Error
	if err != nil {
		return c.RenderError(errors.New("Record Not Found")) // if there is no error, call Render() method to create the HTML file
	}
	return c.Render(posts)
}

func (c Post) Create() revel.Result {
	post := models.Post {
		Body: c.Params.Form.Get("body"),
	}

	res := DB.Create(&post)
	if res.Error != nil {
		return c.RenderError(errors.New("Record Create Failure"+ res.Error.Error()))
	}
	return c.Redirect("/posts")
}

func (c Post) Delete() revel.Result {
	id := c.Params.Route.Get("id") //get the ":id" part of the URL: "/posts/:id/delete"
	posts := []models.Post{}
	res := DB.Delete(&posts, id)
	if res.Error != nil {
		return c.RenderError(errors.New("Record Delete failure. " + res.Error.Error()))
	}
	return c.Redirect("/posts")
}
