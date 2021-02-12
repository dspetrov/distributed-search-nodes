package networking

import (
	"bytes"
	"dspetrov/distributed-search/model"
	"io/ioutil"
	"net/http"

	proto "github.com/gogo/protobuf/proto"
)

type WebClient struct {
}

func NewWebClient() *WebClient {
	wc := WebClient{}
	return &wc
}

func (wc WebClient) SendTask(url string, requestPayload []byte, ch chan model.Result) {
	// To disable connection pooling use:
	// t := http.DefaultTransport.(*http.Transport).Clone()
	// t.DisableKeepAlives = true
	// c := &http.Client{Transport: t}
	// response, err := c.Post(url, "application/octet-stream", bytes.NewReader(requestPayload))

	response, err := http.Post(url, "application/octet-stream", bytes.NewReader(requestPayload))
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	bytes, _ := ioutil.ReadAll(response.Body)

	var result model.Result
	if err := proto.Unmarshal(bytes, &result); err != nil {
		panic(err)
	}

	ch <- result
}
