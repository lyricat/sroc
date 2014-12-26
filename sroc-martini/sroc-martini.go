package main

import (
	"github.com/go-martini/martini"
	"github.com/lyricat/sroc"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"log"
	"net/http"
)

type EchoForm struct {
	Name string `form:"name"`
}

func loadMidares(m *martini.ClassicMartini) {
	m.Use(render.Renderer())
}

func initViews(cors *sroc.CORS, m *martini.ClassicMartini) {
	m.Options("/**", cors.Preflighted)
	m.Post("/echo", cors.AllowCORS, binding.Bind(EchoForm{}), ServPost)
	m.Get("/echo", cors.AllowCORS, ServGet)
	m.Delete("/user/:name", cors.AllowCORS, ServDel)
}

func ServGet(r render.Render, req *http.Request) {
	qs := req.URL.Query()
	vars := make(map[string]interface{})
	vars["Name"] = qs.Get("name")
	r.JSON(200, vars)
}

func ServPost(r render.Render, echo EchoForm) {
	vars := make(map[string]interface{})
	vars["Name"] = echo.Name
	r.JSON(200, vars)
}

func ServDel(r render.Render, params martini.Params) {
	vars := make(map[string]interface{})
	vars["Name"] = params["name"]
	r.JSON(200, vars)
}

func main() {
	cors := sroc.Create(sroc.ANY_ORIGIN, sroc.ALL_METHODS, "X-Authorization, Authorization")
	m := martini.Classic()
	loadMidares(m)
	initViews(cors, m)

	log.Printf("Server Starts...\n")
	log.Fatal(http.ListenAndServe("0.0.0.0:9876", m))
}
