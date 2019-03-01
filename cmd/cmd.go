package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/YuShuanHsieh/trello-transform/server"
	"github.com/YuShuanHsieh/trello-transform/system"
)

func main() {
	filePath := flag.String("p", "", "The path of exported json file from Trello")
	stopSignal := flag.Bool("stop", false, "Stop Web Server")
	flag.Parse()
	if *filePath != "" {
		req := createTransformRequest(*filePath)
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Panicf("[Send request error] %s \n", err.Error())
		}
		data, _ := ioutil.ReadAll(res.Body)
		fmt.Printf("%s", data)
		return
	}
	if *stopSignal {
		req := createStopServerRequest()
		http.DefaultClient.Do(req)
		return
	}
}

func createStopServerRequest() *http.Request {
	port := system.GetPortNumber(server.DefaultPort)
	url := fmt.Sprintf("http://localhost:%d/server/stop", port)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Panicf("[Cannot create a request] %s \n", err.Error())
	}
	return req
}

func createTransformRequest(filePath string) *http.Request {
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	fw, _ := w.CreateFormFile("file", filePath)

	file, err := os.Open(filePath)
	if err != nil {
		log.Panicln("Cannot open the file. Please check the path is correct")
	}
	defer file.Close()

	io.Copy(fw, file)
	w.Close()

	port := system.GetPortNumber(server.DefaultPort)
	url := fmt.Sprintf("http://localhost:%d/transform", port)
	req, err := http.NewRequest("POST", url, &buf)
	if err != nil {
		log.Panicf("[Cannot create a request] %s \n", err.Error())
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	return req
}
