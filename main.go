package main

import (
	"github.com/domesama/Golang-Echo-REST/commands/web"
)

func main()  {
	web.Start()
}






//-------------------------------Example---------------------------------

	//Type Casting
	//http.Handle("/", http.HandlerFunc(defaultRouteHandler))

	//Normal HandleFunc
	//http.HandleFunc("/", defaultRouteHandler)

	//Both are the same

	//Starting the server
	//fmt.Println("Started server on http://localhost:8080")
	//log.Fatal(http.ListenAndServe(":8080", nil))

//----------------------------------------------------------------------



//func defaultRouteHandler(w http.ResponseWriter, r *http.Request){
//	fmt.Fprintf(w,fmt.Sprintf("The current request is + \n %v", r))
//}