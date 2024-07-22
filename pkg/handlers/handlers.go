package handlers

import (
	"net/http"

	"github.com/AMSteffensen/goapp/pkg/render"
)

// Home is the homepage handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}