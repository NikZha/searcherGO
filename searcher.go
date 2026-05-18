package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"sync"
	"time"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	httpLoad := readHttpFile("home.html")
	w.Write(httpLoad)
}

func readHttpFile(filename string) []byte {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Println(err)
	}
	return file
}

type Links struct {
	StrRequest string   `json:"strrequest"`
	Links      []string `json:"links"`
}

func postlinksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Println(w, http.StatusMethodNotAllowed, "Only POST allowed")
		return
	}
	var links Links
	if err := json.NewDecoder(r.Body).Decode(&links); err != nil {
		log.Println(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	defer r.Body.Close()
	var wg sync.WaitGroup
	var mu sync.Mutex
	allResults := make([]ResposJson, 0, len(links.Links))

	for _, url := range links.Links {
		wg.Add(1)

		// Запускаем горутину для каждого URL
		go func(url string) {
			defer wg.Done()

			_, body := getBody(url)
			var respJson ResposJson
			respJson.Emails = getEmail(body)
			respJson.Url = url

			// Безопасно добавляем в общий срез
			mu.Lock()
			allResults = append(allResults, respJson)
			mu.Unlock()
		}(url) // Передаём URL как параметр!
	}

	wg.Wait() // Ждём завершения всех горутин

	// Отправляем результат
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allResults)
}

var client = &http.Client{
	Timeout: 30 * time.Second,
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		if len(via) >= 5 { 
			return http.ErrUseLastResponse
		}
		return nil
	},
}

func getBody(url string) (int, string) {
	resp, err := client.Get(url) 
	if err != nil {
		log.Printf("Error fetching %s: %v\n", url, err)
		return 0, ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading body %s: %v\n", url, err)
		return resp.StatusCode, ""
	}

	return resp.StatusCode, string(body)
}

func getEmail(htmlPage string) []string {
	emailRegex := regexp.MustCompile(`[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}`)
	emails := emailRegex.FindAllString(htmlPage, -1)
	return emails
}

type ResposJson struct {
	Emails []string `json:"emails"`
	Url    string   `json:"url"`
}

func IsPortAvailable(port int) bool {
	addr := fmt.Sprintf(":%d", port)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return false
	}

	listener.Close()

	return true
}

func getPort(portNumber int) int {
	for {
		if IsPortAvailable(portNumber) {
			break
		} else {
			portNumber += 1
		}
	}
	return portNumber
}

func main() {
	portNumber := 9000
	portNumber = getPort(portNumber)
	httpPort := ":" + strconv.Itoa(portNumber)
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/postlinks", postlinksHandler)
	fmt.Println("Сервер запущен на http://localhost" + httpPort)

	// Запускаем сервер
	log.Fatal(http.ListenAndServe(httpPort, nil))
}
