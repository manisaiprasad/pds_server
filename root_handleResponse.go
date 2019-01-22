package main

import (
	"fmt"
	"net/http"
)

type myHandler struct{
	greeting string
}

func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%v World ",mh.greeting)))

}

func main() {
	http.Handle("/", &myHandler{greeting:"Hello"})
	http.ListenAndServe(":8000",nil)




	/*resp, err := http.Get("http://google.com/")

	if err != nil {
		fmt.Println("error",err)
		os.Exit(1)
	}

	fmt.Println(resp,"\n")
	fmt.Printf(resp.Status)*/
}
