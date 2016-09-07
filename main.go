package main

import (
	"fmt"
	"./database"
	"./request"
	"./pod"
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
	db := database.NewDb("root", "123456", "tcp", "localhost:3306", "proxydb")

	//request.GetRequestByPath(db, "/users")

	//re := request.Request{
	//	RequestPath : "abc",
	//	ServiceId:1,
	//	Method    :"GET",
	//	CpuCost   :1,
	//	MemCost    :1,
	//	TimeCost    :1,
	//	Path:"/haha",
	//}
	//request.CreateRequest(db, re)
	//request.DeleteRequest(db, 6)

	requestList := request.GetAllRequest(db)
	for _, r := range requestList {
		fmt.Println(r.Path)
	}
	fmt.Println("-----------------")
	podList := pod.GetAllPod(db)
	for _, p := range podList {
		fmt.Println(p.PodName)
	}
	fmt.Println("-----------------")
	database.CloseDb(db)

	http.HandleFunc("/hello", HelloServer)
	http.ListenAndServe(":12345", nil)
}
