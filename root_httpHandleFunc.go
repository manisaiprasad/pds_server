package main

import (
	"net/http"
)

/*type myHandlerHf struct{
	greeting string
}

func (mh myHandlerHf) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%v World ",mh.greeting)))

}*/

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		writer.Write([]byte("Hello World from handleFunc"))
	})
	http.ListenAndServe(":8000",nil)




	/*resp, err := http.Get("http://google.com/")

	if err != nil {
		fmt.Println("error",err)
		os.Exit(1)
	}

	fmt.Println(resp,"\n")
	fmt.Printf(resp.Status)*/
}
