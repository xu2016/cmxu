package xwb

import (
	"html/template"
	"log"
	"net/http"
)

//HTMLPage 返回Html格式的页面
func HTMLPage(w http.ResponseWriter, name string, t []string, rp interface{}) {
	var tp *template.Template
	var err error
	tp, err = template.ParseFiles(t...)

	if err != nil {
		log.Println("模板加载:", err)
	}
	tp.ExecuteTemplate(w, name, rp)
}
