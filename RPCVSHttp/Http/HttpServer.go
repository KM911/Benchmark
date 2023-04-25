package Http

import (
	"httpvsrpc/config"
	"net/http"
	"strconv"
	"strings"
)

// 消耗一点时间 作用和time.Sleep一样
// 增大时间差异
func CreateString(length int) string {
	builder := strings.Builder{}
	for i := 0; i < length; i++ {
		builder.WriteString("a")
	}
	return builder.String()
}

func HttpServer() {
	// 创建一个http服务器
	server := http.Server{
		Addr: config.HttpServerPort,
	}
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		length := request.Form.Get("length")
		// string to int
		atoi, _ := strconv.Atoi(length)
		writer.Write([]byte(CreateString(atoi)))
	})
	server.ListenAndServe()
}
