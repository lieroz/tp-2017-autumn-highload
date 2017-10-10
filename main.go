package main

import (
	"flag"
	"runtime"
)

func main() {
	port := flag.String("p", ":8080", "a string")
	webRoot := flag.String("wr", "/var/www/html", "a string")
	numCpu := flag.Int("c", runtime.NumCPU(), "an int")
	flag.Parse()

	runtime.GOMAXPROCS(*numCpu)
	serv := NewServer(*port, *webRoot)
	serv.ListenAndServe()
}
