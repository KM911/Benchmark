package Http

import (
	"fmt"
	"httpvsrpc/config"
	"io/ioutil"
	"net/http"
)

func HttpRequest() {
	url := "http://localhost" + config.HttpServerPort + "/hello?length=10"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 这里的开销也有一部分 我们可以将其删除 然后比较其结果 发现影响不大 只有百分之几的结果影响
	//req.Header.Add("User-Agent", "Apifox/1.0.0 (https://www.apifox.cn)")
	//req.Header.Add("Accept", "*/*")
	//req.Header.Add("Host", "localhost:5000")
	//req.Header.Add("Connection", "keep-alive")
	for i := 0; i < 1000; i++ {
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print(string(body)[:0])
	}
}
