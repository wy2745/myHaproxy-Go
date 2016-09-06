package main

import (
	"fmt"
	"./database"
	"./request"
	"net/http"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func HelloServer(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		http.Redirect(w, req, "http://202.120.40.175:21101/users/all", http.StatusTemporaryRedirect)
	} else {
		http.Redirect(w, req, "http://202.120.40.175:21101/users", http.StatusTemporaryRedirect)
	}
}

func main() {
	fmt.Print("haha")
	db := database.NewDb("root", "123456", "tcp", "localhost:3306", "proxydb")
	requestList := request.GetAllRequest(db)
	for _, r := range requestList {
		fmt.Println(r.Path)
	}
	//request.GetRequestByPath(db, "/users")
	database.CloseDb(db)
	
	http.HandleFunc("/hello", HelloServer)
	http.ListenAndServe(":12345", nil)
}
