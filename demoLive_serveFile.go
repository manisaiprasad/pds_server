package main

import "net/http"

func main() {
	/*http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "authy_pds"+r.URL.Path)
	})*/
	http.ListenAndServe(":8000",http.FileServer(http.Dir("authy_pds")))
}
