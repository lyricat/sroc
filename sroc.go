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
}

func Create(origin, methods string) (sroc *CORS) {
	sroc = new(CORS)
	sroc.AllowOrigin = origin
	sroc.AllowMethods = methods
	return sroc
}

func (s *CORS) AllowCORS(res http.ResponseWriter) {
	res.Header().Set("Access-Control-Allow-Origin", s.AllowOrigin)
}

func (s *CORS) Preflighted(req *http.Request, res http.ResponseWriter) {
	res.Header().Set("Access-Control-Allow-Origin", s.AllowOrigin)
	res.Header().Set("Access-Control-Allow-Methods", s.AllowMethods)
	res.Header().Set("Access-Control-Allow-Credentials", "false")
	res.Write(nil)
}
