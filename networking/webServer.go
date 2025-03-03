package networking

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
)

const STATUS_ENDPOINT = "/status"

type WebServer struct {
	port              int
	server            http.Server
	onRequestCallback OnRequestCallback
}

func NewWebServer(port int, onRequestCallback OnRequestCallback) *WebServer {
	ws := WebServer{
		port:              port,
		onRequestCallback: onRequestCallback,
	}

	return &ws
}

func (ws *WebServer) StartServer() {
	m := http.NewServeMux()
	m.HandleFunc(STATUS_ENDPOINT, ws.handleStatusCheckRequest)
	m.HandleFunc(ws.onRequestCallback.GetEndpoint(), ws.handleTaskRequest)

	ws.server = http.Server{Addr: fmt.Sprint(":", ws.port), Handler: m}

	if err := ws.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
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

	fmt.Fprint(w, string(responseBytes))
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

	fmt.Fprint(w, "Server is alive\n")
}

func (ws *WebServer) Stop() {
	ws.server.Shutdown(context.Background())
}
