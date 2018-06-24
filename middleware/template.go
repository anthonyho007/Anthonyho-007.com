package middleware

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

var layoutFunction = template.FuncMap{
	"yield": func() (string, error) {
		return "", fmt.Errorf("yield function failed")
	},
}

var layoutTemplate = template.Must(
	template.New("template.html").
		Funcs(layoutFunction).
		ParseFiles("templates/template.html"),
)

var errorTemplate, _ = ioutil.ReadFile("templates/errorTemplate.html")

var templates = template.Must(
	template.New("templates").
		ParseGlob("templates/**/*.html"),
)

func RenderTemplate(
	writer http.ResponseWriter,
	res *http.Request,
	name string,
	data map[string]interface{},
) {

	if data == nil {
		data = map[string]interface{}{}
	}

	cloneLayout, _ := layoutTemplate.Clone()
	cloneLayout.Funcs(template.FuncMap{
		"yield": func() (template.HTML, error) {
			buf := bytes.NewBuffer(nil)
			err := templates.ExecuteTemplate(buf, name, data)
			return template.HTML(buf.String()), err
		},
	})
	err := cloneLayout.Execute(writer, data)
	if err != nil {
		printError(writer, name, err)
	}
}

func printError(writer http.ResponseWriter, name string, err error) {
	http.Error(
		writer,
		fmt.Sprintf(string(errorTemplate), name, err),
		http.StatusInternalServerError,
	)
}
