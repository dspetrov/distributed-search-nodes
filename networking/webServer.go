package networking

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const STATUS_ENDPOINT = "/status"

type WebServer struct {
	port              int
	onRequestCallback OnRequestCallback
}

func NewWebServer(port int, onRequestCallback OnRequestCallback) WebServer {
	ws := WebServer{
		port:              port,
		onRequestCallback: onRequestCallback,
	}

	return ws
}

func (ws *WebServer) StartServer() {
	http.HandleFunc(STATUS_ENDPOINT, ws.handleStatusCheckRequest)
	http.HandleFunc(ws.onRequestCallback.GetEndpoint(), ws.handleTaskRequest)

	if err := http.ListenAndServe(fmt.Sprint(":", ws.port), nil); err != nil {
		panic(err)
	}
}

func (ws *WebServer) handleTaskRequest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != ws.onRequestCallback.GetEndpoint() {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	requestBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	responseBytes := ws.onRequestCallback.HandleRequest(requestBytes)

	fmt.Fprintf(w, string(responseBytes))
}

func (ws *WebServer) handleStatusCheckRequest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != STATUS_ENDPOINT {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Server is alive\n")
}
