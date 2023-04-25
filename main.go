package main

import (
	"flag"
	"log"
)

// Log前缀设定
func init() {
	log.SetPrefix("Bench_Management: ")
}

func main() {
	port := flag.Uint("port", 5000, "TCP Port Number for Bench Management Server")
	gateway := flag.String("gateway", "http://127.0.0.1:5000", "Bench Management Gateway")
	flag.Parse()
	app := NewBenchManagementServer(uint16(*port), *gateway)
	app.Run()
}
