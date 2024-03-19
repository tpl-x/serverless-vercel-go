package api

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<h1>Hello Vercel!</h1>")
}
