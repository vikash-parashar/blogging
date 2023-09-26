package render

import (
	"fmt"
	"html/template"
	"io"
	"log"
)

func RenderTemplates(w io.Writer, tmpName string, data interface{}) {
	tmp, err := template.ParseFiles(fmt.Sprintf("./templates/%s.html", tmpName))
	if err != nil {
		log.Println(err)
	}
	err = tmp.Execute(w, data)
	if err != nil {
		log.Println(err)
	}
}
