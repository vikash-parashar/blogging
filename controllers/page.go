package controllers

import (
	"blogging/render"

	"github.com/gin-gonic/gin"
)

func RenderHomepage(c *gin.Context) {
	render.RenderTemplates(c.Writer, "home", nil)
}

func RenderLoginPage(c *gin.Context) {
	render.RenderTemplates(c.Writer, "login", nil)
}

func RenderRegisterPage(c *gin.Context) {
	render.RenderTemplates(c.Writer, "register", nil)
}
