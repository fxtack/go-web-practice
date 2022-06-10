package middleware

import "net/http"

type CrossMiddleWare struct {
	Next http.Handler
}

func (c CrossMiddleWare) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if c.Next == nil {
		c.Next = http.DefaultServeMux
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	c.Next.ServeHTTP(w, r)
}
