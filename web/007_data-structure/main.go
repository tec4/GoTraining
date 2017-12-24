package main

import (
	"text/template"
	"os"
	"log"
	"strings"
)

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
}

func init() {
	 tpl = template.Must(template.New("").Funcs(fm).ParseGlob("*.gohtml"))
}

func main() {
	sages := []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad"}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}
}
