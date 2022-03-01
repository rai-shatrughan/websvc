package handler

import (
	// "bytes"
	// "fmt"
	// "io/ioutil"
	"log"
	"net/http"
	// "time"
)

type Transport struct {
	// Uncomment this if you want to capture the transport
	// CapturedTransport http.RoundTripper
}

type Montioringpath struct {
	Path        string
	Count       int64
	Duration    int64
	DurationMS  string
	AverageTime int64
}

var globalMap = make(map[string]Montioringpath)

func (t *Transport) RoundTrip(request *http.Request) (*http.Response, error) {

	log.Println("-----------New Request--------------")
	log.Println("Request URI : ", request.RequestURI)
	log.Println("Request Method : ", request.Method)

	// if request.Body != nil {
	// 	buf, _ := ioutil.ReadAll(request.Body)
	// 	// rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
	// 	rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))

	// 	// log.Println("Request body : ", rdr1)
	// 	request.Body = rdr2 // OK since rdr2 implements the
	// }

	// start := time.Now()
	response, err := http.DefaultTransport.RoundTrip(request)
	if err != nil {
		log.Println("Error Processing Request : ", err)
		return nil, err //Server is not reachable. Server not working
	}
	// elapsed := time.Since(start)

	// key := request.Method + "-" + request.URL.Path //for example for POST Method with /path1 as url path key=POST-/path1

	// if val, ok := globalMap[key]; ok {
	// 	val.Count = val.Count + 1
	// 	val.Duration += elapsed.Milliseconds()
	// 	val.DurationMS = fmt.Sprint(elapsed)
	// 	val.AverageTime = val.Duration / val.Count
	// 	globalMap[key] = val
	// 	//do something here
	// } else {
	// 	var m Montioringpath
	// 	m.Path = request.URL.Path
	// 	m.Count = 1
	// 	m.Duration = elapsed.Milliseconds()
	// 	m.AverageTime = m.Duration / m.Count
	// 	globalMap[key] = m
	// }
	// b, err := json.MarshalIndent(globalMap, "", "  ")
	// if err != nil {
	// 	fmt.Println("error:", err)
	// }

	// body, err := httputil.DumpResponse(response, true)
	// if err != nil {
	// 	print("\n\nerror in dumb response")
	// 	// copying the response body did not work
	// 	return nil, err
	// }

	// log.Println("Response Body : ", string(body))
	// log.Println("Request URI : ", request.RequestURI)
	// log.Println("Request Method : ", request.Method)
	// log.Println("Response Code : ", response.StatusCode)
	// log.Println("Response Time : ", elapsed)

	// log.Printf("Path : %s - Count : %d - Response Time : %s - Avg. Response Time : %dms", globalMap[key].Path, globalMap[key].Count, globalMap[key].DurationMS, globalMap[key].AverageTime)

	return response, err
}
