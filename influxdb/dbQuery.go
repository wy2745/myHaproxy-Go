package influxdb

import (
	"bytes"
	"fmt"
	"net/http"
	"io/ioutil"
	"../json"
	"net/url"
)

const (
	cpu_rateUrl string = "https://sjtu.caicloudapp.com/api/v1/proxy/namespaces/kube-system/services/monitoring-influxdb:api/query?db=k8s&q=SELECT \"usage_rate\" FROM \"cpu\" WHERE \"type\" = 'pod_container' AND \"namespace_name\" = 'kube-system' AND \"pod_name\" = 'elasticsearch-logging-v1-82mud'"
	influxdbQuery string = "https://sjtu.caicloudapp.com/api/v1/proxy/namespaces/kube-system/services/monitoring-influxdb:api/query?"
	limitNum string = " limit 1"
	usageRateNamespacePodName string = "SELECT \"usage_rate\" FROM \"cpu\" WHERE \"type\" = 'pod_container' AND \"namespace_name\" = 'kube-system' AND \"pod_name\" = "
	dbstr string = "db"
	db string = "k8s"
	qstr string = "q"
)

func generateQueryUrl(desurl string, podName string) string {
	var value url.Values
	value = make(url.Values)
	value.Add(dbstr, db)
	value.Add(qstr, usageRateNamespacePodName + "'" + podName + "'" + limitNum)
	str := value.Encode()
	fmt.Println(desurl + str)
	return desurl + str
}

func setBasicAuthOfCaicloud(r *http.Request) {
	r.SetBasicAuth(userName, password)
}

func InvokeRequest_Caicloud(method string, ourl string, body []byte) *http.Response {

	client := http.Client{}
	var req *http.Request
	var err error
	if body != nil {
		req, err = http.NewRequest(method, generateQueryUrl(ourl, "elasticsearch-logging-v1-82mud"), bytes.NewBuffer(body))
	} else {
		req, err = http.NewRequest(method, generateQueryUrl(ourl, "elasticsearch-logging-v1-82mud"), nil)
	}

	setBasicAuthOfCaicloud(req)
	if method != "GET" {
		req.Header.Set("Content-Type", "application/json")
	}
	resp, err := client.Do(req)
	//fmt.Println(resp.Header)
	//fmt.Println(resp.Status)
	//fmt.Println(resp.StatusCode)
	if err != nil {
		fmt.Print(err)
	}
	return resp
}

func InvokeGetReuqest(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Print(err)
		return nil
	}
	return resp
}

func GetCpuRate(mode string) {
	var resp *http.Response
	if mode == Test {
		resp = InvokeGetReuqest(influxdbQuery)
	} else {
		resp = InvokeRequest_Caicloud("GET", influxdbQuery, nil)

	}
	if (resp != nil) {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if (err != nil) {
			fmt.Print(err)
			return
		}
		var v results
		jsonParse.JsonUnmarsha(body, &v)
		fmt.Println(v)
		//for _, item := range v.Items {
		//	classType.PrintNode(item)
		//}
	}
}
