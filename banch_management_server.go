package main

import (
	"log"
	"net/http"
	"path"
	"text/template"
)

const tempDir = "./templates"

type DemoManagementServer struct {
	port    uint16
	gateway string
}

func NewBenchManagementServer(port uint16, gateway string) *DemoManagementServer {
	return &DemoManagementServer{port, gateway}
}

func (dm *DemoManagementServer) Port() uint16 {
	return dm.port
}

func (dm *DemoManagementServer) Gateway() string {
	return dm.gateway
}

// http.HandleFunc("/", dm.Index)
func (dm *DemoManagementServer) Index(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		t, _ := template.ParseFiles(path.Join(tempDir, "index.html"))
		t.Execute(w, "")
	default:
		log.Printf("ERROR: Invalid HTTP Method")
	}
}

// http.HandleFunc("/wallet", ws.Wallet)
func (dm *DemoManagementServer) Demo(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		w.Header().Add("Content-Type", "application/json")
		//myDemo := NewDemo()
		//fmt.Println(myDemo)
		//m, _ := myDemo.MarshalJSON()
		//io.WriteString(w, string(m[:]))
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: Invalid HTTP Method")
	}
}

func (dm *DemoManagementServer) Run() error {
	return nil
}
