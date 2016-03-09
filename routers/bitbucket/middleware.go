package bitbucket

import (
	"fmt"
	"net/http"
)

func NormalizePRPayload(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		handler.ServeHTTP(w, r)
	})
}
