package backend

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/api/helloword", helloworld)
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}