package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

var servers map[string]func() = make(map[string]func())

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	servers["http"] = bootstrapHttp
	servers["websocket"] = bootstrapWebsocket
	servers["default"], _ = servers["http"]
	//servers["default"], _ = servers["websocket"]

	server := "default"

	if len(os.Args) == 2 {
		if serverTypeIdx := strings.Index(os.Args[1], "--"); serverTypeIdx != -1 {
			server = os.Args[1][(serverTypeIdx + 2):]
		}
	}

	if serverFunc, ok := servers[server]; ok {
		serverFunc()
	} else {
		panic(fmt.Sprintf("No server has been found for %s", server))
	}
}
