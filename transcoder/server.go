package transcoder

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

type Server *martini.ClassicMartini

func NewServer(session *DatabaseSession) Server {
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Use(session.Database())
	setupRoutes(m)
	return m
}

func setupRoutes(m Server) {
	m.Get("/", Home)
	m.Post("/api/actions", binding.Json(File{}), GetActions)
}
