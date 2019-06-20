package controllers

import (
	"github.com/revel/revel"
	// "golang.org/x/crypto/bcrypt"
	"webapp/app/routes"
	// "webapp/app/models"
	

)

type Auth struct {
	BaseController
}

func (c Auth) Login() revel.Result {
	return c.Render()
}

func (c Auth) checkUser() revel.Result {
	return c.Render()
}
// func (c AuthController) Authenticate(username, password string, remember bool) revel.Result {
// 	user := c.getUser(username)

// 	if user != nil {
// 		err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
// 		if err == nil {
// 			c.Session["user"] = username
// 			if remember {
// 				c.Session.SetDefaultExpiration()

// 			}else{
// 				c.Session.SetNoExpiration()
// 			}
// 			c.Flash.Success("welcome, " + username)
// 			return c.Redirect(routes.App.Index())
// 		}
// 	}

// 	c.Flash.Out["username"] = username
// 	c.Flash.Error("login failed")
// 	return c.Redirect(routes.App.Index())
// }

// func (c AuthController) getUser(username string) (user *models.User) {
// 	user = &models.User{}
// 	_, err := c.Session.GetInto("fulluser", user, false)
// 	if user.Username == username {
// 		return user
// 	}
// 	return models.FindOrCreate("guest")
// }


func (c Auth) logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(routes.App.Index())
}