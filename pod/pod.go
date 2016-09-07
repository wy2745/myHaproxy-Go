package pod

import (
	"fmt"
	"database/sql"
	"strings"
)

type Pod struct {
	PodName    string
	CpuUsage   float32
	MemUsage   float32
	Address    string
	ServiceId  int
	Ability    string
	Connection int
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func scanPod(rows *sql.Rows) Pod {
	var pod Pod
	err := rows.Scan(&pod.PodName, &pod.CpuUsage, &pod.MemUsage,
		&pod.Address, &pod.ServiceId, &pod.Ability, &pod.Connection)
	checkErr(err)
	fmt.Println("Connection:", pod.Connection)
	fmt.Println("PodName:", pod.PodName)
	fmt.Println("ServiceId:", pod.ServiceId)
	fmt.Println("Address:", pod.Address)
	fmt.Println("CpuUsage", pod.CpuUsage)
	fmt.Println("MemUsage:", pod.MemUsage)
	fmt.Println("Ability:", pod.Ability)
	fmt.Println("-----------------")
	return pod
}
func GetAllPod(db *sql.DB) []Pod {
	var podList []Pod
	rows, err := db.Query("SELECT * FROM Pod")
	checkErr(err)
	for rows.Next() {
		podList = append(podList, scanPod(rows))
	}
	return podList
}
func escapeMysqlQuery(path string) string {
	str := strings.Replace(path, "/", "\"/", 1)
	return str + "\""
}

func GetPodByName(db *sql.DB, podName string) Pod {
	var pod Pod
	rows, err := db.Query("select * From Pod where Pod.podName = " + podName)
	checkErr(err)
	for rows.Next() {
		pod = scanPod(rows)
	}
	return pod
}

func CreatePod(db *sql.DB, pod Pod) Pod {
	stmt, err := db.Prepare("INSERT Pod SET podName=?,cpuUsage=?,memUsage=?,address=?,serviceId=?,ability=?,connection=?")
	checkErr(err)

	res, err := stmt.Exec(pod.PodName, pod.CpuUsage, pod.MemUsage, escapeMysqlQuery(pod.Address), pod.ServiceId, pod.Ability, pod.Connection)
	checkErr(err)
	fmt.Println(res.RowsAffected())
	return pod
}

func DeleteRequest(db *sql.DB, podName string) {
	stmt, err := db.Prepare("delete from Pod where podName=?")
	checkErr(err)
	res, err := stmt.Exec(podName)
	checkErr(err)
	fmt.Println(res.RowsAffected())

}

