package gotcha

import(
  "fmt"
  "net/http"
)

type HTTPErroCodeSignaller interface {
  Raise(w http.ResponseWriter)
}

type Gotcha struct {
  Message string
  ErrorCode int
}

func(p Gotcha) Raise(w http.ResponseWriter)(){
  w.Header().Set("Content-Type", "application/json; charset=utf-8")
  w.WriteHeader(p.ErrorCode)
  fmt.Fprintf(w, p.Message)
}
