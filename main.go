package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

const port = 80

var (
	index    string
	tplIndex string
)

type indexData struct {
	Hostname string
}

func init() {
	tplIndex = `
<html>
  <head>
    <title>Hello, Docker!</title>
    <style>
      body {
        text-align: center;
        font-family: "Open Sans","Helvetica Neue",Helvetica,Arial,sans-serif;
      }
    </style>
  </head>
  <body>
    <div class="container">
      <img src="data:image/png;base64,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"></img>
      <h1>Hello, container world!</h1>
      <p>My hostname is <b>{{ .Hostname }}</b></p>
    </div>
  </body>
</html>
`
}

func main() {
	listen := fmt.Sprintf(":%d", port)

	http.HandleFunc("/", indexHandler)

	log.Printf("starting server on port %d", port)
	log.Fatal(http.ListenAndServe(listen, nil))
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Fprintf(res, err.Error())
	}

	data := indexData{Hostname: hostname}

	tpl, err := template.New("index").Parse(tplIndex)
	if err != nil {
		fmt.Fprintf(res, err.Error())
	}

	tpl.Execute(res, data)
}
