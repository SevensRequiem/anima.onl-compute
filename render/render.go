package render

import (
	"github.com/gin-contrib/multitemplate"
)

func Create() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("index", "views/base.html", "views/index.html")
	return r
}
