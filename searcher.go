package main

import (
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

func main() {
	portNumber := 9000
	httpPort := ":" + strconv.Itoa(portNumber)
	http.HandleFunc("/", homeHandler)
	fmt.Println("Сервер запущен на http://localhost:9000")

	// Запускаем сервер
	log.Fatal(http.ListenAndServe(httpPort, nil))
}
