package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	gretting := "This is sanjok"
	return c.Render(gretting)
}

func (c App) Hello(myName string) revel.Result{
	c.Validation.Required(myName).Message("Name required")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}
	return c.Render(myName)
}
