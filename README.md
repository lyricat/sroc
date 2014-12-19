# SROC

SROC is a simple package for writing CORS services in Golang.

## Getting Started

Create CORS handler.

```
    cors := sroc.Create(sroc.ANY_ORIGIN, sroc.ALL_METHODS)
```

Adding CORS support to the service

```
func initViews(cors *sroc.CORS, m *martini.ClassicMartini) {
    m.Options("/**", cors.Preflighted)
    m.Post("/echo", cors.AllowCORS, binding.Bind(EchoForm{}), ServPost)
    m.Get("/echo", cors.AllowCORS, ServGet)
    m.Delete("/user/:name", cors.AllowCORS, ServDel)
}
```


