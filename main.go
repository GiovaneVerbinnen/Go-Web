package main

import (
	"net/http"
	"text/template"

	"github.com/GiovaneVerbinnen/Go-Web/routes"
	_ "github.com/lib/pq"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
