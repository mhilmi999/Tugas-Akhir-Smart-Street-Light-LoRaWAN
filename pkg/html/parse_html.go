package html

import (
	"log"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
)

func Render(templatesDir string) multitemplate.Renderer{
	r := multitemplate.NewRenderer()
	layouts, err := filepath.Glob(templatesDir + "layouts/*.tmpl")
	
	if err != nil{
		log.Println(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/**/*")

	for _, include := range includes{
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}

func ManualRender(tmpDir string) multitemplate.Renderer{
	r := multitemplate.NewRenderer()
	
	return r
}