package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	httpLoad := readHttpFile("home.html")
	w.Write(httpLoad)
}

func readHttpFile(filename string) []byte {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

type Links struct {
	StrRequest string   `json:"strrequest"`
	Links      []string `json:"links"`
}

func postlinksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Fatalln(w, http.StatusMethodNotAllowed, "Only POST allowed")
		return
	}
	var links Links
	if err := json.NewDecoder(r.Body).Decode(&links); err != nil {
		log.Fatalln(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()
	fmt.Println("Get links:", links)
}

func main() {
	portNumber := 9000
	httpPort := ":" + strconv.Itoa(portNumber)
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/postlinks", postlinksHandler)
	fmt.Println("Сервер запущен на http://localhost:9000")

	// Запускаем сервер
	log.Fatal(http.ListenAndServe(httpPort, nil))
}
