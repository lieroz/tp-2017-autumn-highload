package main

import (
	"flag"
)

func main() {
	port := flag.String("p", ":8080", "a string")
	webRoot := flag.String("wr", "/static", "a string")
	flag.Parse()
	serv := NewServer(*port, *webRoot)
	serv.ListenAndServe()
}
