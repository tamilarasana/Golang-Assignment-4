package main

import (
	  
	  "log"
	  "net/http"
	  "fmt"
	)



func  main(){
	http.HandleFunc("/users", userHandleFunc)
	log.Fatal(http.ListenAndServe(":8080",nil))
}

func userHandleFunc(w http.ResponseWriter, r *http.Request){
 fmt.Println("we got a request on /users")
  fmt.Println(w, "hi thanks my user api /users api http method")

// userCounter++
// s := fmt.Sprint(userApiRes, r.Method userCounter)
// fmt.fPrint(w,s)

}