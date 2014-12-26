package sroc

import (
	"net/http"
)

const (
	ANY_ORIGIN  = "*"
	ALL_METHODS = "GET, POST, PUT, DELETE"
)

type CORS struct {
	AllowOrigin  string
	AllowMethods string
	AllowHeaders string
}

func Create(origin, methods, headers string) (sroc *CORS) {
	sroc = new(CORS)
	sroc.AllowOrigin = origin
	sroc.AllowMethods = methods
	sroc.AllowHeaders = headers
	return sroc
}

func (s *CORS) AllowCORS(res http.ResponseWriter) {
	res.Header().Set("Access-Control-Allow-Origin", s.AllowOrigin)
}

func (s *CORS) Preflighted(res http.ResponseWriter) {
	res.Header().Set("Access-Control-Allow-Origin", s.AllowOrigin)
	res.Header().Set("Access-Control-Allow-Methods", s.AllowMethods)
	res.Header().Set("Access-Control-Allow-Headers", s.AllowHeaders)
	res.Header().Set("Access-Control-Allow-Credentials", "false")
	res.Write(nil)
}
