package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/student", ActionStudent)

	server := &http.Server{
		Addr:    ":9000",
		Handler: nil,
	}

	fmt.Println("Server started at localhost:9000")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if !Auth(w, r) {
		return
	}
	if !AllowOnlyGet(w, r) {
		return
	}

	if id := r.URL.Query().Get("id"); id != "" {
		OutputJson(w, SelectStudent(id))
		return
	}

	OutputJson(w, GetStudents())
}

func OutputJson(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
