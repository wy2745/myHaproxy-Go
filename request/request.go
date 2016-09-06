package request

import (
	"database/sql"
	"fmt"
	"strings"
)

type Request struct {
	RequestId   int
	RequestPath string
	ServiceId   int
	Method      string
	CpuCost     float32
	MemCost     float32
	TimeCost    float32
	Path        string
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func scanRequest(rows *sql.Rows) Request {
	var request Request
	err := rows.Scan(&request.RequestId, &request.RequestPath, &request.ServiceId,
		&request.Method, &request.CpuCost, &request.MemCost, &request.TimeCost, &request.Path)
	checkErr(err)
	fmt.Println("requestId:", request.RequestId)
	fmt.Println("requestRequestPath:", request.RequestPath)
	fmt.Println("requestServiceId:", request.ServiceId)
	fmt.Println("requestMethod:", request.Method)
	fmt.Println("requestCpuCost", request.CpuCost)
	fmt.Println("requestMemCost:", request.MemCost)
	fmt.Println("requestTimeCost:", request.TimeCost)
	fmt.Println("requestPath:", request.Path)
	fmt.Println("-----------------")
	return request
}
//TODO:预计加上返回查询结果的数组
func GetAllRequest(db *sql.DB) {

	rows, err := db.Query("SELECT * FROM Request")
	checkErr(err)
	for rows.Next() {
		scanRequest(rows)
	}
}
func escapeMysqlQuery(path string) string {
	str := strings.Replace(path, "/", "\"/", 1)
	return str + "\""
}

func GetRequestByPath(db *sql.DB, path string) Request {
	var request Request
	str := escapeMysqlQuery(path)
	rows, err := db.Query("select * From Request where Request.path = " + str)
	checkErr(err)
	for rows.Next() {
		request = scanRequest(rows)
	}
	return request
}
