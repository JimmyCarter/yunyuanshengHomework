package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	os.Setenv("VERSION", "1.2")
	log.SetFlags(log.Lmicroseconds | log.Ldate)
	// httpPandler := wrapHandlerWithLogging()
	log.SetPrefix("[云原生]")
	fmt.Printf("%s", os.Getenv("VERSION"))
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/", rootHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("%s", err)
	}
}
func rootHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Printf("%T", v)
		w.Header().Add(k, fmt.Sprintf("%s", v))
	}
	statusCode := 200
	w.WriteHeader(statusCode)
	log.Printf("IP=%s\nHTTP return code=%s\n", r.RemoteAddr, http.StatusText(statusCode))
	w.Header().Add("VERSION", os.Getenv("VERSION"))
	resp := make(map[string]string)
	resp["message"] = "success"
	resp["status"] = "200"
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
func healthz(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "success"
	resp["status"] = "200"
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
