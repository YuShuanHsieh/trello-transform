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
)

func main() {
	filePath := flag.String("p", "trello.json", "The path of exported json file from Trello")
	if *filePath == "" {
		log.Panicln("The file path cannot be empty")
	}

	flag.Parse()

	req := createTransformRequest(*filePath)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Panicf("[Send request error] %s \n", err.Error())
	}
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("%s", data)
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

	req, err := http.NewRequest("POST", "http://localhost:8181/transform", &buf)
	if err != nil {
		log.Panicf("[Cannot create a request] %s \n", err.Error())
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	return req
}
