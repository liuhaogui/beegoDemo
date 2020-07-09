package controllers

import (
	"github.com/astaxie/beego"
	"github.com/liuhaogui/beegoDemo/utils"
)

// UserAuth
type AuthController struct {
	beego.Controller
}

// @Title Login
// @Description Login
// @Param   email     query   string  false       "邮箱"
// @Param   password     query   string  false       "密码"
// @Success 200
// @Failure 400
// @router /login [post]
func (c *AuthController) Login() {
	c.Data["json"] = utils.JsonMsg("", utils.SUCCESS, "")
	c.ServeJSON()
	return
}
