package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
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
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/proxydb?charset=utf8")
	checkErr(err)
	rows, err := db.Query("SELECT * FROM Pod")
	checkErr(err)
	for rows.Next() {
		var podName string
		var cpuUsage float32
		var memUsage float32
		var address string
		var serviceId int
		var ability float32
		var connection int
		err = rows.Scan(&podName, &cpuUsage, &memUsage, &address, &serviceId, &ability, &connection)
		checkErr(err)
		fmt.Println(podName)
		fmt.Println(cpuUsage)
		fmt.Println(memUsage)
		fmt.Println(address)
		fmt.Println(serviceId)
		fmt.Println(ability)
		fmt.Println(connection)
	}
	http.HandleFunc("/hello", HelloServer)
	http.ListenAndServe(":12345", nil)
}
